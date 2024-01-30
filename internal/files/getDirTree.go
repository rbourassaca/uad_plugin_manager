package files

import (
	"os"
	"path/filepath"
	"rbourassa/uadPluginManager/internal/config"
	"strings"
)

func GetDirTree(path string) []string {
	var files []string
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for i := 0; i < len(config.Config.PluginFormats); i++ {
			if strings.Contains(info.Name(), config.Config.PluginFormats[i].Extension) && !strings.Contains(path, "Contents") {
				files = append(files, path)
			}
		}
		return nil
	})
	return files
}
