package config

import (
	"log"
	"os/user"
)

type Config struct {
	PluginsPath              string
	VimawesomePluginUrl      string
	VimawesomePluginQueryUrl string
	Messages                 map[string]string
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
	messages["provide_plugin_name"] = "please, provide a plugin name" + eol
	messages["request_error"] = "there was an error with the request" + eol
	messages["parse_error"] = "there was an error parsing the response" + eol

	config := &Config{
		PluginsPath:              usr.HomeDir + "/.vim/bundle/",
		VimawesomePluginUrl:      "http://vimawesome.com/api/plugins/%s",
		VimawesomePluginQueryUrl: "http://vimawesome.com/api/plugins?query=%s",
		Messages:                 messages,
	}

	return config
}
