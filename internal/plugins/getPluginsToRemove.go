package plugins

import (
	"fmt"
	"rbourassa/uadPluginManager/internal/config"
	"strings"
)

func GetPluginsToRemove(collectionsName []string) []string {
	var pluginsToRemove []string
	for i := range len(collectionsName) { // For each collections to remove
		collection := config.Config.PluginDefinition[strings.ToLower(collectionsName[i])]
		if len(collection) > 0 {
			for x := range len(collection) { // For each plugins in the current collection
				pluginsToRemove = append(pluginsToRemove, collection[x])
			}
		} else {
			fmt.Println("Unable to find collection " + collectionsName[i])
		}
	}
	return pluginsToRemove
}
