package plugins

import (
	"fmt"
	"rbourassa/uadPluginManager/internal/config"
	"strings"
)

func GetPluginsToRemove(collectionsName []string) []string {
	var pluginsToRemove []string
	for i := 0; i < len(collectionsName); i++ { // For each collections to remove
		collection := config.Config.PluginDefinition[strings.ToLower(collectionsName[i])]
		if len(collection) > 0 {
			for x := 0; x < len(collection); x++ { // For each plugins in the current collection
				pluginsToRemove = append(pluginsToRemove, collection[x])
			}
		} else {
			fmt.Println("Unable to find collection " + collectionsName[i])
		}
	}
	return pluginsToRemove
}
