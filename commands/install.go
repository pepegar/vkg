package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pepegar/vkg-go/config"
)

func IsUserRepo(path string) bool {
	match, _ := regexp.MatchString("^[a-z\\-\\.]*\\/[a-z\\-\\.]*$", path)

	return match
}

func IsGithubUrl(path string) bool {
	match, _ := regexp.MatchString("^github.com\\/[a-z\\-\\.]*\\/[a-z\\-\\.]*$", path)

	return match
}

func IsVimawesomeSlug(param string) bool {
	return (!IsGithubUrl(param) && !IsUserRepo(param))
}

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
			var slug string
			var url string

			param := os.Args[2]

			if IsUserRepo(param) {
				parts := strings.Split(param, "/")
				slug = parts[len(parts)-1]
				url = "https://github.com/" + param
			} else if IsGithubUrl(param) {
				parts := strings.Split(param, "/")
				slug = parts[len(parts)-1]
				url = "https://" + param
			} else if IsVimawesomeSlug(param) {
				jsonUrl := fmt.Sprintf(config.VimawesomePluginUrl, param)
				body, requestError := GetJson(jsonUrl)

				if requestError != nil {
					log.Fatal(config.Messages["request_error"])
				} else {
					plugin, parseError := ParseSinglePlugin(body)

					if parseError != nil {
						log.Fatal(parseError)
					} else {
						url = plugin.GithubUrl
						slug = plugin.Slug
					}
				}
			}

			Install(url, slug)
		}
	},
}
