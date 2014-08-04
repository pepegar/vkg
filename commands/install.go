package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/pepegar/vkg-go/config"
)

func Install(url string, name string) bool {
	config := config.GetVkgGonfig()

	cmd := exec.Command("git", "clone", url+".git", config.PluginsPath+name)
	_, err := cmd.Output()

	if err != nil {
		fmt.Printf(config.Messages["plugin_already_installed"], name)
		return false
	}

	fmt.Printf(config.Messages["successfully_installed"], name)
	return true
}

var InstallCommand = Command{
	Name:        "install",
	Description: "Installs a package from vimawesome",
	Usage:       "install <package>",
	Action: func() {
		config := config.GetVkgGonfig()
		if len(os.Args) < 3 {
			log.Fatal(config.Messages["provide_plugin_name"])
		} else {
			url := fmt.Sprintf(config.VimawesomePluginUrl, os.Args[2])
			body, requestError := GetJson(url)

			if requestError != nil {
				log.Fatal(config.Messages["request_error"])
			} else {
				plugin, parseError := ParseSinglePlugin(body)

				if parseError != nil {
					log.Fatal(parseError)
				} else {
					Install(plugin.GithubUrl, plugin.Slug)
				}
			}
		}
	},
}
