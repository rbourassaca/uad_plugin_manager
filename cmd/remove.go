package cmd

import (
	"fmt"
	"rbourassa/uadPluginManager/internal/config"
	"rbourassa/uadPluginManager/internal/file"
	"strings"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove plugins that aren't owned",
	Run: func(cmd *cobra.Command, args []string) {
		UADSystemProfile := file.Open(config.Config.Files.UADSystemProfile)
		CollectionsToRemove := file.Find(UADSystemProfile, ": Demo not started", true)
		for i := 0; i<len(CollectionsToRemove); i++ { // For each collections to remove
			collection := config.Config.PluginDefinition[strings.ToLower(CollectionsToRemove[i])]
			for x := 0; x < len(collection); x++ { // For each plugins in the current collection
				plugin := collection[x]
				fmt.Println(plugin)
			}
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
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
