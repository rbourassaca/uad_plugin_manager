package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"rbourassa/uadPluginManager/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "uad_plugin_manager",
	Short: "Simple tool to manage UAD plugins",
	Long:  "UAD Plugin Manager is a tool used to manage Universal Audio DSP powered plugins.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config.Load()
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.uadPluginManager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
