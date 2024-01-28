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
		UADSystemProfile := file.Open(viper.Get("files.UADSystemProfile").(string))
		toRemove := file.Find(UADSystemProfile, ": Demo not started", true)
		fmt.Println(toRemove)
		// Remove plugins
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
