package main

import (
	"fmt"
	"rbourassa/uadPluginManager/cmd"
	"runtime"
	"slices"
)

func main() {
	if slices.Contains([]string{"windows", "darwin"}, runtime.GOOS) {
		cmd.Execute()
	} else {
		fmt.Println("Incompatible with current OS.")
	}
}
