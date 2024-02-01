package cmd

import (
	"fmt"
	"rbourassa/uadPluginManager/internal/config"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List collections",
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		if all {
			for collection, plugins := range config.Config.PluginDefinition {
				fmt.Println("\n" + collection)
				for _, value := range plugins {
					fmt.Println("|--> " + strings.ToLower(value))
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().BoolP("all", "a", false, "List all collections")
}
