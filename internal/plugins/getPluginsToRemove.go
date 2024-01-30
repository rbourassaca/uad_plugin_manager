package plugins

import (
	"rbourassa/uadPluginManager/internal/config"
	"strings"
)

func GetPluginsToRemove(collections []string) []string {
	var pluginsToRemove []string
	for i := 0; i < len(collections); i++ { // For each collections to remove
		collection := config.Config.PluginDefinition[strings.ToLower(collections[i])]
		for x := 0; x < len(config.Config.PluginDefinition[strings.ToLower(collections[i])]); x++ { // For each plugins in the current collection
			pluginsToRemove = append(pluginsToRemove, collection[x])
		}
	}
	return pluginsToRemove
}
