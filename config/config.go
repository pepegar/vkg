package config

import (
	"log"
	"os"
	"os/user"
)

type Config struct {
	PluginsPath              string
	VkgrcPath                string
	VimawesomePluginUrl      string
	VimawesomePluginQueryUrl string
	Messages                 map[string]string
}

func VkgrcExists() bool {
	if _, err := os.Stat(GetVkgGonfig().VkgrcPath); err == nil {
		return true
	}

	return false
}

func GetVkgGonfig() *Config {
	usr, errUser := user.Current()

	if errUser != nil {
		log.Fatal(errUser)
	}

	eol := "\r\n"

	messages := make(map[string]string)
	messages["plugin_already_installed"] = "plugin %s is already installed" + eol
	messages["successfully_installed"] = "%s successfully installed" + eol
	messages["successfully_deleted"] = "%s successfully deleted" + eol
	messages["provide_plugin_name"] = "please, provide a plugin name" + eol
	messages["request_error"] = "there was an error with the request" + eol
	messages["parse_error"] = "there was an error parsing the response" + eol
	messages["plugin_does_not_exist"] = "plugin %s does not exist" + eol

	config := &Config{
		PluginsPath:              usr.HomeDir + "/.vim/bundle/",
		VkgrcPath:                usr.HomeDir + "/.vkgrc",
		VimawesomePluginUrl:      "http://vimawesome.com/api/plugins/%s",
		VimawesomePluginQueryUrl: "http://vimawesome.com/api/plugins?query=%s",
		Messages:                 messages,
	}

	return config
}
