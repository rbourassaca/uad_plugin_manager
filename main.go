package main

import (
	"rbourassa/uadPluginManager/cmd"
	"rbourassa/uadPluginManager/internal/config"
)

func main() {
	config.Load()
	cmd.Execute()
}
