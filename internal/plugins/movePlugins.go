package plugins

import (
	"fmt"
	"path/filepath"
	"rbourassa/uadPluginManager/internal/config"
	"rbourassa/uadPluginManager/internal/files"
	"strings"
)

func MovePlugins(pluginsToMove []string) {
	for i := 0; i < len(config.Config.PluginFormats); i++ {
		glob, _ := filepath.Glob(config.Config.PluginFormats[i].Path + "/*" + config.Config.PluginFormats[i].Extension)
		if len(glob) == 0 {
			glob, _ = filepath.Glob(config.Config.PluginFormats[i].Path + "/**/*" + config.Config.PluginFormats[i].Extension)
		}
		pluginFormatName := config.Config.PluginFormats[i].Name
		for x := 0; x < len(pluginsToMove); x++ {
			currentPath := GetPluginPaths(pluginsToMove[x], glob)
			for _, currentGlob := range glob {
				if currentGlob == currentPath {
					pluginLocation := strings.TrimPrefix(currentPath, config.Config.PluginFormats[i].Path)
					newPath := filepath.Join(config.RemovedPluginDir, pluginFormatName, pluginLocation)
					err := files.Move(currentPath, newPath)
					if err != nil {
						fmt.Println(err)
					}
				}
			}

		}
	}
}
