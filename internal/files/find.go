package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Find(name string, directory []string) (string, error) {
	for i := 0; i < len(directory); i++ {
		path := filepath.Join(directory[i], name)
		_, err := os.Stat(path)
		if !errors.Is(err, os.ErrNotExist) {
			return path, nil
		}
	}
	return "", fmt.Errorf("unable to find %s in %s", name, directory)
}
