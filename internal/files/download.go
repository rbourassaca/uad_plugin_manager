package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Download(location string, url string) error {
	// Setting the file location
	location = filepath.Join(location, filepath.Base(url))

	// Getting to data from the URL
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(res io.ReadCloser) {
		err := res.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}(res.Body)

	// Creating the file
	file, err := os.Create(location)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}(file)

	// Write data to file
	_, err = io.Copy(file, res.Body)
	return err
}
