package files

import (
	"os"
	"path/filepath"
)

func Move(currentPath string, newPath string) error {
	err := os.MkdirAll(filepath.Dir(newPath), os.FileMode(0522))
	if err == nil {
		err = os.Rename(currentPath, newPath)
	}
	return err
}
