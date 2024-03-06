package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"rbourassa/uadPluginManager/internal/files"
	"runtime"

	"github.com/spf13/viper"
)

type (
	Type struct {
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

var Config Type
var Appdata string
var RemovedPluginDir string
var gitRepository = "https://raw.githubusercontent.com/rbourassaca/uad_plugin_manager/main/"
var configFileName = "config." + runtime.GOOS + ".yaml"
var pluginDefinitionFileName = "pluginDefinition.yaml"

func Load() {
	initEnv()
	handleConfigFiles()
	cleanPaths()
}

func initEnv() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		os.Exit(1)
	}
	Appdata = filepath.Join(userConfigDir, "/UAD-Plugin-Manager")
	RemovedPluginDir = filepath.Join(Appdata, "/Plugins")
	err = os.Mkdir(Appdata, os.ModePerm)
	if err != nil {
		if !errors.Is(err, fs.ErrExist) {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func handleConfigFiles() {
	viper.AddConfigPath(filepath.Clean("./"))
	viper.AddConfigPath(Appdata)
	viper.SetConfigType("yaml")

	loadConfigFile(configFileName)
	loadConfigFile(pluginDefinitionFileName)

	err := viper.UnmarshalExact(&Config)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cleanPaths() {
	Config.Files.UADSystemProfile = filepath.Clean(Config.Files.UADSystemProfile)
	for i := 0; i < len(Config.PluginFormats); i++ {
		Config.PluginFormats[i].Path = filepath.Clean(Config.PluginFormats[i].Path)
	}
}

func loadConfigFile(name string) {
	viper.SetConfigName(name)
	err := viper.ReadInConfig()
	if err != nil {
		if errors.Is(err, err.(viper.ConfigFileNotFoundError)) {
			err := files.Download(Appdata, gitRepository+name)
			if err == nil {
				loadConfigFile(name)
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
