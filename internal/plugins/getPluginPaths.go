package plugins

import (
	"strings"
)

func GetPluginPaths(fileName string, tree []string) string {
	for i := 0; i < len(tree); i++ {
		if strings.Contains(tree[i], fileName) {
			return tree[i]
		}
	}
	return fileName
}
