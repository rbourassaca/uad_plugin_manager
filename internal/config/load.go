package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type (
	ConfigType struct {
		Files            FilesType
		PluginFormats    []PluginFormatType
		PluginDefinition map[string][]string
	}

	FilesType struct {
		UADSystemProfile string
	}

	PluginFormatType struct {
		Path      string
		Extension string
		Name      string
	}
)

var Config ConfigType
var Appdata string
var RemovedPluginDir string

func Load() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		os.Exit(1)
	}
	Appdata = filepath.Join(userConfigDir, "/UAD-Plugin-Manager")
	RemovedPluginDir = filepath.Join(Appdata, "/Plugins")
	os.Mkdir(Appdata, os.ModePerm)
	if runtime.GOOS == "windows" {
		viper.SetConfigFile(filepath.FromSlash("./configs/config.win.yaml"))
	} else if runtime.GOOS == "darwin" {
		viper.SetConfigFile(filepath.FromSlash("./configs/config.osx.yaml"))
	}
	viper.ReadInConfig()
	viper.SetConfigFile(filepath.FromSlash("./configs/pluginDefinition.yaml"))
	viper.MergeInConfig()
	viper.UnmarshalExact(&Config)
	clean()
}

func clean() {
	Config.Files.UADSystemProfile = filepath.Clean(Config.Files.UADSystemProfile)
	for i := 0; i < len(Config.PluginFormats); i++ {
		Config.PluginFormats[i].Path = filepath.Clean(Config.PluginFormats[i].Path)
	}
}
