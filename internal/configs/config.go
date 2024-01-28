package configs

import (
	"github.com/spf13/viper"
)


func LoadConfig()  {
	viper.SetConfigFile("config.toml")
	viper.ReadInConfig()
}