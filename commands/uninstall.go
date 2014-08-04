package commands

import (
	"fmt"
	"os"

	"github.com/pepegar/vkg-go/config"
)

var UninstallCommand = Command{
	Name:        "uninstall",
	Description: "uninstall a local package",
	Usage:       "uninstall <name>",
	Action: func() {
		config := config.GetVkgGonfig()
		if len(os.Args) < 3 {
			fmt.Println(config.Messages["provide_plugin_name"])
		}
	},
}
