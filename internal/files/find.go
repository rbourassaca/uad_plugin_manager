package files

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Find(file *os.File, text []string, crop bool) []string {
	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for i := 0; i < len(text); i++ {
			if strings.Contains(scanner.Text(), text[i]) {
				line := scanner.Text()
				if crop {
					line = strings.Split(line, text[i])[0]
				}
				list = append(list, line)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}
