package plugins

import (
	"fmt"
	"rbourassa/uadPluginManager/internal/config"
	"strings"
)

func ListAllPlugins() {
	for collection, plugins := range config.Config.PluginDefinition {
		fmt.Println("\n" + collection)
		for _, value := range plugins {
			fmt.Println("|--> " + strings.ToLower(value))
		}
	}
}
