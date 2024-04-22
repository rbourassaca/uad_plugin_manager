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
	Short: "Update configurations to the latest one available on github",
	Run: func(cmd *cobra.Command, args []string) {
		configuration, _ := cmd.Flags().GetBool("configuration")
		pluginDefinition, _ := cmd.Flags().GetBool("pluginDefinition")

		if configuration {
			os.Remove(filepath.Join(config.Appdata, config.ConfigFileName))
			files.Download(config.Appdata, config.GitRepository+config.ConfigFileName)
		}
		if pluginDefinition {
			os.Remove(filepath.Join(config.Appdata, config.PluginDefinitionFileName))
			files.Download(config.Appdata, config.GitRepository+config.PluginDefinitionFileName)
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
	updateCmd.Flags().BoolP("configuration", "c", false, "Update paths configuration.")
	updateCmd.Flags().BoolP("pluginDefinition", "p", false, "Update plugins definition")
}
