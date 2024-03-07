package plugins

import (
	"fmt"
	"path/filepath"
	"rbourassa/uadPluginManager/internal/config"
	"rbourassa/uadPluginManager/internal/files"
	"slices"
	"strconv"
	"strings"
	"time"
)

func MovePlugins(pluginsToMove []string) {
	currentUnixTime := strconv.FormatInt(time.Now().Unix(), 10)
	for i := 0; i < len(config.Config.PluginFormats); i++ {
		glob, _ := filepath.Glob(config.Config.PluginFormats[i].Path + "/*" + config.Config.PluginFormats[i].Extension)
		if len(glob) == 0 {
			glob, _ = filepath.Glob(config.Config.PluginFormats[i].Path + "/**/*" + config.Config.PluginFormats[i].Extension)
		}
		for x := 0; x < len(pluginsToMove); x++ {
			currentPath := GetPluginPaths(pluginsToMove[x], config.Config.PluginFormats[i].Extension, glob)
			if slices.Contains(glob, currentPath) {
				pluginLocation := strings.TrimPrefix(currentPath, config.Config.PluginFormats[i].Path)
				newPath := filepath.Join(config.RemovedPluginDir, currentUnixTime, config.Config.PluginFormats[i].Name, pluginLocation)
				err := files.Move(currentPath, newPath)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}
}
