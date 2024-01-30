package files

import (
	"log"
	"os"
)

func Open(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
