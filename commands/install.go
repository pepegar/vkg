package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/pepegar/vkg/config"
	"github.com/pepegar/vkg/config/vkgrc"
	"github.com/pepegar/vkg/utils"
)

func IsUserRepo(path string) bool {
	match, _ := regexp.MatchString("^[a-z\\-\\.]+\\/[a-z\\-\\.]+$", path)

	return match
}

func IsGithubUrl(path string) bool {
	match, _ := regexp.MatchString("^github.com\\/[a-z\\-\\.]+\\/[a-z\\-\\.]+$", path)

	return match
}

func IsVimawesomeSlug(plugin string) bool {
	return (!IsGithubUrl(plugin) && !IsUserRepo(plugin))
}

func installAllVkgrcPlugins() {
	vkgConfig := config.GetVkgGonfig()
	vkgrcContents, err := ioutil.ReadFile(vkgConfig.VkgrcPath)

	if err != nil {
		log.Fatal(err)
	}

	vkgrc := vkgrc.ParseVkgrc(vkgrcContents)

	for _, plugin := range vkgrc.Plugins {
		if err := utils.Git.Clone(plugin.Repository, plugin.Branch); err == nil {
			fmt.Printf(vkgConfig.Messages["successfully_installed"], plugin.Repository)
		} else {
			fmt.Printf(vkgConfig.Messages["plugin_already_installed"], plugin.Repository)
		}
	}
}

func installSinglePlugin(plugin string) {
	vkgConfig := config.GetVkgGonfig()
	var slug string
	var url string

	if IsUserRepo(plugin) {
		parts := strings.Split(plugin, "/")
		slug = parts[len(parts)-1]
		url = "https://github.com/" + plugin
	} else if IsGithubUrl(plugin) {
		parts := strings.Split(plugin, "/")
		slug = parts[len(parts)-1]
		url = "https://" + plugin
	} else if IsVimawesomeSlug(plugin) {
		jsonUrl := fmt.Sprintf(vkgConfig.VimawesomePluginUrl, plugin)
		body, requestError := GetJson(jsonUrl)

		if requestError != nil {
			log.Fatal(vkgConfig.Messages["request_error"])
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

	if err := utils.Git.Clone(url, "master"); err == nil {
		fmt.Printf(vkgConfig.Messages["successfully_installed"], slug)
	} else {
		fmt.Println(err)
	}
}

func actionInstall() {
	if len(os.Args) < 3 && config.VkgrcExists() {
		installAllVkgrcPlugins()
	} else {
		plugin := os.Args[2]
		installSinglePlugin(plugin)
	}
}

var InstallCommand = Command{
	Name:        "install",
	Description: "Installs a package from vimawesome",
	Usage:       "install <package>",
	Action:      actionInstall,
}
