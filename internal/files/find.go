package files

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Find(file *os.File, text string, crop bool) []string {
	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), text) {
			line := scanner.Text()
			if crop {
				line = strings.Split(line, text)[0]
			}
			list = append(list, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}
