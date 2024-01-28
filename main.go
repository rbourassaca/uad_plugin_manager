package main

import (
	"rbourassa/uadPluginManager/cmd"
	"rbourassa/uadPluginManager/internal/configs"
)

func main() {
	configs.LoadConfig()
	cmd.Execute()
}
