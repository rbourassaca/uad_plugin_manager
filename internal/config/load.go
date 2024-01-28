package config

import (
	"github.com/spf13/viper"
)


func Load()  {
	viper.SetConfigFile("config.toml")
	viper.ReadInConfig()
}