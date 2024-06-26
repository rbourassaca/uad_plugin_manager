package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"rbourassa/uadPluginManager/internal/files"
	"runtime"

	"github.com/spf13/viper"
)

type (
	Type struct {
		Files            FilesType
		PluginFormats    []PluginFormatType
		PluginDefinition map[string][]string
	}

	FilesType struct {
		UADSystemProfile string
	}

	PluginFormatType struct {
		Path      string
		Extension string
		Name      string
	}
)

var Config Type
var Appdata string
var RemovedPluginDir string
var GitRepository = "https://raw.githubusercontent.com/rbourassaca/uad_plugin_manager/main/"
var ConfigFileName = "config." + runtime.GOOS + ".yaml"
var PluginDefinitionFileName = "pluginDefinition.yaml"

func Load() {
	initEnv()
	handleConfigFiles()
	cleanPaths()
}

func initEnv() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		os.Exit(1)
	}
	Appdata = filepath.Join(userConfigDir, "/UAD-Plugin-Manager")
	RemovedPluginDir = filepath.Join(Appdata, "/Plugins")
	err = os.Mkdir(Appdata, os.ModePerm)
	if err != nil {
		if !errors.Is(err, fs.ErrExist) {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func handleConfigFiles() {
	viper.AddConfigPath(filepath.Clean("./"))
	viper.AddConfigPath(Appdata)
	viper.SetConfigType("yaml")

	loadConfigFile(ConfigFileName)
	loadConfigFile(PluginDefinitionFileName)

	err := viper.UnmarshalExact(&Config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Config.Files.UADSystemProfile = searchUADSystemProfile()
}

func cleanPaths() {
	for i := 0; i < len(Config.PluginFormats); i++ {
		Config.PluginFormats[i].Path = filepath.Clean(Config.PluginFormats[i].Path)
	}
}

func loadConfigFile(name string) {
	viper.SetConfigName(name)
	err := viper.MergeInConfig()
	if err != nil {
		if errors.Is(err, err.(viper.ConfigFileNotFoundError)) {
			err := files.Download(Appdata, GitRepository+name)
			if err == nil {
				loadConfigFile(name)
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func searchUADSystemProfile() string {
	path, err := files.Find("UADSystemProfile.txt", []string{"./", Appdata})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return path
}
