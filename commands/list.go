package commands

import (
	"fmt"
	"io/ioutil"

	"github.com/pepegar/vkg/config"
)

var ListCommand = Command{
	Name:        "list",
	Description: "list all installed packages",
	Usage:       "list",
	Action: func() {
		config := config.GetVkgGonfig()

		files, _ := ioutil.ReadDir(config.PluginsPath)

		for _, file := range files {
			fmt.Println("* " + file.Name())
		}
	},
}
