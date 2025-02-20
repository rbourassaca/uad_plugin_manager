package config

import (
	"fmt"
	"os"
	"path/filepath"
	"rbourassa/uadPluginManager/internal/files"
	"runtime"

	"github.com/spf13/viper"
)

type (
	Type struct {
		UserData         string
		Repository       string
		Files            FilesType
		PluginFormats    Os
		PluginDefinition map[string][]string
	}

	FilesType struct {
		UADSystemProfile string
	}

	Os struct {
		Windows []PluginFormatType
		Darwin  []PluginFormatType
	}

	PluginFormatType struct {
		Path      string
		Extension string
		Name      string
	}
)

var Config Type

const Runtime = runtime.GOOS

func Load() {
	handleConfigFiles()
}

func handleConfigFiles() {
	viper.AddConfigPath(filepath.Clean("./config"))
	viper.SetConfigType("yaml")
	loadConfigFile("config.yaml")
	loadConfigFile("pluginDefinition.yaml")
	err := viper.UnmarshalExact(&Config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Config.Files.UADSystemProfile = searchUADSystemProfile()
}

func loadConfigFile(name string) {
	viper.SetConfigName(name)
	err := viper.MergeInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func searchUADSystemProfile() string {
	path, err := files.Find("UADSystemProfile.txt", []string{"./config"})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return path
}
