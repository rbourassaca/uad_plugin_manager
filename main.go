package main

import (
	"fmt"
	"rbourassa/uadPluginManager/internal/configs"
	"runtime"

	"github.com/TwiN/go-color"
	"github.com/inancgumus/screen"
)

func main() {
	// Setting up display
	resetScreen()

	// Loading config
	configs.LoadConfig()

	// Checking plugins path
	if(runtime.GOOS == "windows"){
		// if windows
		fmt.Println("OS			Windows")
		fmt.Println("UADSystemProfile	" + configs.Config.Files.UADSystemProfile)
		fmt.Println("VST2			" + "\"" + configs.Config.Path["win"].VST2 + "\"")
		fmt.Println("VST3			" + "\"" + configs.Config.Path["win"].VST3 + "\"")
		fmt.Println("AAX			" + "\"" + configs.Config.Path["win"].AAX + "\"")
	} 

	// Confirm
	fmt.Println("\nPress enter to continue...")
	fmt.Scanln()
}

func resetScreen() {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Println(color.InBlackOverGray("UAD Plugin Manager")+"\n")
}
