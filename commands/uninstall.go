package commands

import (
	"fmt"
	"os"

	"github.com/pepegar/vkg/config"
)

func uninstallAction(config config.Config) {
	if len(os.Args) < 3 {
		fmt.Println(config.Messages["provide_plugin_name"])
	} else {
		pluginName := os.Args[2]
		pluginPath := config.PluginsPath + pluginName
		if _, err := os.Stat(pluginPath); err == nil {
			os.RemoveAll(pluginPath)
			fmt.Printf(config.Messages["successfully_deleted"], pluginName)
		} else {
			fmt.Printf(config.Messages["plugin_does_not_exist"], pluginName)
		}
	}
}

var UninstallCommand = Command{
	Name:        "uninstall",
	Description: "uninstall a local package",
	Usage:       "uninstall <name>",
	Action:      uninstallAction,
}
