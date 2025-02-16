package plugins

import (
	"os"
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
		currentPluginFormatPath := filepath.Clean(pluginFormats[i].Path)
		filepath.Walk(currentPluginFormatPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if match, _ := filepath.Match("*"+pluginFormats[i].Extension, filepath.Base(path)); match {
				if slices.Contains(pluginsToMove, strings.TrimSuffix(filepath.Base(path), filepath.Ext(pluginFormats[i].Extension))) {
					newPath := filepath.Join(config.Config.UserData, currentUnixTime, pluginFormats[i].Name, strings.TrimPrefix(path, currentPluginFormatPath))
					files.Move(path, newPath)
				}

				if info.IsDir() {
					return filepath.SkipDir
				}
			}
			return nil
		})
	}
}
