package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"

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

func getPlugin(plugin vkgrc.VkgrcPlugin, wg *sync.WaitGroup, vkgConfig *config.Config) {
	defer wg.Done()

	cloneError := utils.Git.Clone(plugin.Repository, plugin.Branch)

	if cloneError == nil {
		fmt.Printf(vkgConfig.Messages["successfully_installed"], plugin.Repository)
	} else {
		fmt.Printf(vkgConfig.Messages["plugin_already_installed"], plugin.Repository)
	}
}

func installAllVkgrcPlugins() {
	var wg sync.WaitGroup
	vkgConfig := config.GetVkgGonfig()
	vkgrcContents, err := ioutil.ReadFile(vkgConfig.VkgrcPath)

	if err != nil {
		log.Fatal(err)
	}

	parsedVkgrc := vkgrc.ParseVkgrc(vkgrcContents)

	for _, plugin := range parsedVkgrc.Plugins {
		wg.Add(1)
		go getPlugin(plugin, &wg, vkgConfig)
	}

	wg.Wait()

	fmt.Println("all plugins installed")
}

type Plugin interface {
	GetSlug() string
	GetURL() string
}

type GithubPlugin struct {
	URL  string
	Slug string
}

func (p GithubPlugin) GetSlug() string {
	return p.Slug
}

func (p GithubPlugin) GetURL() string {
	return p.URL
}

type VimawesomePlugin struct {
	URL  string
	Slug string
}

func (p VimawesomePlugin) GetSlug() string {
	return p.Slug
}

func (p VimawesomePlugin) GetURL() string {
	return p.URL
}

func newVimawesomePlugin(name string) VimawesomePlugin {
	var plugin VimawesomePlugin
	vkgConfig := config.GetVkgGonfig()
	jsonUrl := fmt.Sprintf(vkgConfig.VimawesomePluginUrl, name)
	body, requestError := GetJson(jsonUrl)

	if requestError != nil {
		log.Fatal(vkgConfig.Messages["request_error"])
	} else {
		plugRecord, parseError := ParseSinglePlugin(body)

		if parseError != nil {
			log.Fatal(parseError)
		} else {
			plugin = VimawesomePlugin{
				URL:  plugRecord.GithubUrl,
				Slug: plugRecord.Slug,
			}
		}
	}

	return plugin
}

func newGithubPlugin(name string) GithubPlugin {
	parts := strings.Split(name, "/")

	plugin := GithubPlugin{
		Slug: parts[len(parts)-1],
		URL:  "https://" + name,
	}

	return plugin
}

func newUserRepoPlugin(name string) GithubPlugin {
	parts := strings.Split(name, "/")
	plugin := GithubPlugin{
		Slug: parts[len(parts)-1],
		URL:  "https://github.com/" + name,
	}
	return plugin
}

func newPlugin(name string) Plugin {
	var plugin Plugin

	if IsUserRepo(name) {
		plugin = newUserRepoPlugin(name)
	} else if IsGithubUrl(name) {
		plugin = newGithubPlugin(name)
	} else if IsVimawesomeSlug(name) {
		plugin = newVimawesomePlugin(name)
	}

	return plugin
}

func installSinglePlugin(param string) {
	vkgConfig := config.GetVkgGonfig()

	plugin := newPlugin(param)

	if err := utils.Git.Clone(plugin.GetURL(), "master"); err == nil {
		fmt.Printf(vkgConfig.Messages["successfully_installed"], plugin.GetSlug())
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
