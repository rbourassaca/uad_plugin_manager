package plugins

import (
	"fmt"
	"path/filepath"
	"rbourassa/uadPluginManager/internal/config"
	"rbourassa/uadPluginManager/internal/files"
	"strings"
)

func MovePlugins(pluginsToRemove []string) {
	for i := 0; i < len(config.Config.PluginFormats); i++ {
		glob, _ := filepath.Glob(config.Config.PluginFormats[i].Path + "/*" + config.Config.PluginFormats[i].Extension)
		if len(glob) == 0 {
			glob, _ = filepath.Glob(config.Config.PluginFormats[i].Path + "/**/*" + config.Config.PluginFormats[i].Extension)
		}
		pluginFormatName := config.Config.PluginFormats[i].Name
		for x := 0; x < len(pluginsToRemove); x++ {
			currentPath := GetPluginPaths(pluginsToRemove[x], glob)
			pluginLocation := strings.TrimPrefix(currentPath, filepath.FromSlash(config.Config.PluginFormats[i].Path))
			newPath := filepath.Join(config.Appdata, config.RemovedPluginDir, pluginFormatName, pluginLocation)
			err := files.Move(currentPath, newPath)
			fmt.Println(err)
		}
	}
}
