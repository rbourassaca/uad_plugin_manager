package config

import (
	"github.com/spf13/viper"
)

type(
	ConfigType struct{
		Files FilesType
		PluginFormats []PluginFormatType
		PluginDefinition map[string][]string
	}

	FilesType struct {
		UADSystemProfile string;
	}

	PluginFormatType struct {
		Path string
		Extension string
	}
)

var Config ConfigType

func Load()  {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
	viper.SetConfigFile("pluginDefinition.yaml")
	viper.MergeInConfig()
	viper.UnmarshalExact(&Config)
}