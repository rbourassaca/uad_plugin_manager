package main

import (
	"fmt"
	"rbourassa/uadPluginManager/cmd"
	"rbourassa/uadPluginManager/internal/config"
	"runtime"
	"slices"
)

func main() {
	if slices.Contains([]string{"windows", "darwin"}, runtime.GOOS) {
		config.Load()
		cmd.Execute()
	} else {
		fmt.Println("Incompatible with current OS.")
	}
}
