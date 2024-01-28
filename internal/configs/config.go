package configs

import (
	"fmt"
	"os"

	"rbourassa/uadPluginManager/internal/models"

	"github.com/BurntSushi/toml"
)

const configFile string = "./config.toml"

var Config models.Config

func LoadConfig()  {
	_, err := toml.DecodeFile(configFile, &Config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}