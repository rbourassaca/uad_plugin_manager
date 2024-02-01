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
		collectionsName, _ := cmd.Flags().GetStringArray("collectionsName")
		unlicensed, _ := cmd.Flags().GetBool("unlicensed")
		if unlicensed {
			fmt.Println("Removing all unlicensed collections")
			UADSystemProfile := files.Open(config.Config.Files.UADSystemProfile)
			pluginsToRemove := files.Find(UADSystemProfile, []string{": Demo not started"}, true) //TODO: add what is written in UADSystemProfile when demo is over, will update when I have a demo that's over
			unlicensedPlugins := plugins.GetPluginsToRemove(pluginsToRemove)
			plugins.MovePlugins(unlicensedPlugins)
		}
		if len(collectionsName) > 0 {
			fmt.Println("Removing all specified collections")
			plugins.MovePlugins(plugins.GetPluginsToRemove(collectionsName))
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
	removeCmd.Flags().BoolP("unlicensed", "u", false, "Select unlicensed collections")
	removeCmd.Flags().StringArrayP("collectionsName", "n", []string{}, "Select specific collections to remove")
}
