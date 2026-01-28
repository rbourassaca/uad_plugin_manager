package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Find(name string, directory string) (string, error) {
	path := filepath.Join(directory, name)
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		return path, nil
	}
	absoluteDir, _ := filepath.Abs(directory)
	return "", fmt.Errorf("unable to find %s in %v", name, absoluteDir)
}
