package cmd

import (
	"os"
	"path/filepath"
	"rbourassa/uadPluginManager/internal/config"
	"rbourassa/uadPluginManager/internal/files"

	"github.com/spf13/cobra"
)

// updateCmd represents the list command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update selected configuration to the latest one available on GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		configuration, _ := cmd.Flags().GetBool("configuration")
		pluginDefinition, _ := cmd.Flags().GetBool("pluginDefinition")

		if configuration {
			os.Remove(filepath.Join("./config", "config.yaml"))
			files.Download("./config", config.Config.Repository+"/config/config.yaml")
		}
		if pluginDefinition {
			os.Remove(filepath.Join("./config", "pluginDefinition.yaml"))
			files.Download("./config", config.Config.Repository+"/config/pluginDefinition.yaml")
		}
	},
}

// err := files.Download(Appdata, gitRepository+name)

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	updateCmd.Flags().BoolP("configuration", "c", false, "Update configuration.")
	updateCmd.Flags().BoolP("pluginDefinition", "p", false, "Update plugins definition")
}
