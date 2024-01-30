package plugins

import (
	"fmt"
	"rbourassa/uadPluginManager/internal/config"
	"rbourassa/uadPluginManager/internal/files"
	"strings"
)

func MovePlugins(pluginsToRemove []string) {
	for i := 0; i < len(config.Config.PluginFormats); i++ {
		dirTree := files.GetDirTree(config.Config.PluginFormats[i].Path)
		for x := 0; x < len(pluginsToRemove); x++ {
			currentPath := GetPluginPaths(pluginsToRemove[x], dirTree)
			newPath := config.Appdata + config.RemovedPluginDir + "/" + config.Config.PluginFormats[i].Extension[1:] + strings.TrimPrefix(currentPath, config.Config.PluginFormats[i].Path)
			err := files.Move(currentPath, newPath)
			fmt.Println(err)
		}
	}
}
