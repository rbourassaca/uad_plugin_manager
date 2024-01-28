package cmd

import (
	"fmt"
	"rbourassa/uadPluginManager/internal/file"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove plugins that aren't owned",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the UADSystemProfile
		UADSystemProfile := file.Open(viper.Get("files.UADSystemProfile").(string))
		// Load plugins to remove
		toRemove := file.Find(UADSystemProfile, ": Demo not started", true)
		// For each collections to remove
		for i := 0; i<len(toRemove); i++ { 
			// Creating a variable for the current collection
			var collection []string
			viper.UnmarshalKey(fmt.Sprintf("pluginDefinition.%s", toRemove[i]), &collection)
			// For each plugins in the current collection
			for x := 0; x < len(collection); x++ {
				plugin := collection[x]
				fmt.Println(plugin)
				//TODO: Move the current plugin
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
