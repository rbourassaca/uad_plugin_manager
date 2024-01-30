package files

import (
	"os"
)

func Move(oldPath string, newPath string) error {
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		if err = os.MkdirAll(newPath, os.ModePerm); err != nil {
			return err
		}
	}

	return os.Rename(oldPath, newPath)
}
