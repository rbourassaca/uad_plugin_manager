package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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
	viper.AddConfigPath(Appdata)
	viper.AddConfigPath(filepath.Clean("./"))
	viper.SetConfigType("yaml")

	viper.SetConfigName("config." + runtime.GOOS)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigName("pluginDefinition")
	err = viper.MergeInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = viper.UnmarshalExact(&Config)
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
