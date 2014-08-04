package commands

import (
	"fmt"
	"os"

	"github.com/pepegar/vkg-go/config"
)

var SearchCommand = Command{
	Name:        "search",
	Description: "Search a plugin",
	Usage:       "search <plugin>",
	Action: func() {
		config := config.GetVkgGonfig()
		if len(os.Args) < 3 {
			fmt.Println(config.Messages["provide_plugin_name"])
		} else {
			url := fmt.Sprintf(config.VimawesomePluginQueryUrl, os.Args[2])
			json, jsonError := GetJson(url)

			if nil != jsonError {
				fmt.Println(config.Messages["request_error"])
			}

			response, parseJsonError := ParsePluginsList(json)

			if nil != parseJsonError {
				fmt.Println(config.Messages["parse_error"])
			}

			for _, plugin := range response.Plugins {
				fmt.Println("* " + plugin.Slug + " - " + plugin.ShortDesc)
			}
		}
	},
}
