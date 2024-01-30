package config

import (
	"os"

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

const RemovedPluginDir = "/Plugins"

func Load() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		os.Exit(1)
	}
	Appdata = userConfigDir + "/UAD-Plugin-Manager"
	os.Mkdir(Appdata, os.ModePerm)
	viper.SetConfigFile("./configs/config.yaml")
	viper.ReadInConfig()
	viper.SetConfigFile("./configs/pluginDefinition.yaml")
	viper.MergeInConfig()
	viper.UnmarshalExact(&Config)
}
