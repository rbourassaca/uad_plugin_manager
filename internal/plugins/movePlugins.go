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
	pluginFormats := config.Config.PluginFormats.Darwin
	if config.Runtime == "windows" {
		pluginFormats = config.Config.PluginFormats.Windows
	}

	for i := 0; i < len(pluginFormats); i++ {
		glob, _ := filepath.Glob(pluginFormats[i].Path + "/*" + pluginFormats[i].Extension)
		if len(glob) == 0 {
			glob, _ = filepath.Glob(pluginFormats[i].Path + "/**/*" + pluginFormats[i].Extension)
		}
		for x := 0; x < len(pluginsToMove); x++ {
			currentPath := GetPluginPaths(pluginsToMove[x], pluginFormats[i].Extension, glob)
			if slices.Contains(glob, currentPath) {
				pluginLocation := strings.TrimPrefix(currentPath, pluginFormats[i].Path)
				newPath := filepath.Join(config.Config.UserData, "removedPlugins", currentUnixTime, pluginFormats[i].Name, pluginLocation)
				err := files.Move(currentPath, newPath)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}
}
