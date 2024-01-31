package plugins

import (
	"path/filepath"
	"strings"
)

func GetPluginPaths(fileName string, extension string, tree []string) string {
	for i := 0; i < len(tree); i++ {
		if strings.Contains(filepath.Base(tree[i]), fileName+extension) {
			return tree[i]
		}
	}
	return ""
}
