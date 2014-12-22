package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pepegar/vkg/config"
)

func getPlugins() []os.FileInfo {
	config := config.GetVkgGonfig()

	files, _ := ioutil.ReadDir(config.PluginsPath)

	return files
}

func getListOutput(files []os.FileInfo) string {
	var result string

	for _, file := range files {
		result += "* " + file.Name() + "\n"
	}

	return strings.TrimSpace(result)
}

func action() {
	files := getPlugins()

	output := getListOutput(files)

	fmt.Println(output)
}

var ListCommand = Command{
	Name:        "list",
	Description: "list all installed packages",
	Usage:       "list",
	Action:      action,
}
