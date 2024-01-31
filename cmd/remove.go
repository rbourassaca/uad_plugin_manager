package cmd

import (
	"fmt"
	"rbourassa/uadPluginManager/internal/config"
	"rbourassa/uadPluginManager/internal/files"
	"rbourassa/uadPluginManager/internal/plugins"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove plugins.",
	Run: func(cmd *cobra.Command, args []string) {
		unlicensed, _ := cmd.Flags().GetBool("unlicensed")
		if unlicensed {
			unlicensedPlugins := plugins.GetPluginsToRemove(files.Find(files.Open(config.Config.Files.UADSystemProfile), ": Demo not started", true))
			plugins.MovePlugins(unlicensedPlugins)
		} else {
			fmt.Println("No plugins were selected.")
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	removeCmd.Flags().BoolP("unlicensed", "u", false, "Select unlicensed plugins")
}
