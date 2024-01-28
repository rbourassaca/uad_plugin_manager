package config

import (
	"github.com/spf13/viper"
)

type(
	config struct{
		Files files
		Paths paths
		Extensions extensions
		PluginDefinition map[string][]interface{}
	}

	files struct {
		UADSystemProfile string;
	}

	paths struct {
		Win win
		Osx osx
	}

	win struct {
		VST2 string
		VST3 string
		AAX string
	}

	osx struct {
		VST2 string
		VST3 string
		AAX string
		AU string
	}

	extensions struct {
		VST2 string
		VST3 string
		AAX string
		AU string
	}
)

var Config config

func Load()  {
	viper.SetConfigFile("config.toml")
	viper.ReadInConfig()
	viper.SetConfigFile("pluginDefinition.toml")
	viper.MergeInConfig()
	viper.Unmarshal(&Config)
}