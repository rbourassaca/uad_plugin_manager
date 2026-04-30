package plugins

import (
	"path/filepath"
	"strings"
)

func GetPluginPaths(fileName string, extension string, tree []string) string {
	for i := range len(tree) {
		if strings.Contains(filepath.Base(tree[i]), fileName+extension) {
			return tree[i]
		}
	}
	return ""
}
