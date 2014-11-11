package commands

import (
	"fmt"
	"os"

	"github.com/pepegar/vkg/config"
	"github.com/pepegar/vkg/utils"
)

func searchAction(config config.Config) {
	if len(os.Args) < 3 {
		fmt.Println(config.Messages["provide_plugin_name"])
	} else {
		url := fmt.Sprintf(config.VimawesomePluginQueryUrl, os.Args[2])
		json, jsonError := utils.GetJson(url)

		if nil != jsonError {
			fmt.Println(config.Messages["request_error"])
		}

		response, parseJsonError := utils.ParsePluginsList(json)

		if nil != parseJsonError {
			fmt.Println(config.Messages["parse_error"])
		}

		for _, plugin := range response.Plugins {
			fmt.Println("* " + plugin.Slug + " - " + plugin.ShortDesc)
		}
	}
}

var SearchCommand = Command{
	Name:        "search",
	Description: "Search a plugin",
	Usage:       "search <plugin>",
	Action:      searchAction,
}
