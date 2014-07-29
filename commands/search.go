package commands

import (
	"fmt"
	"os"
)

var SearchCommand = Command{
	Name:        "search",
	Description: "Search a plugin",
	Action: func() {
		if len(os.Args) < 3 {
			fmt.Println("please, supply a keyword for the plugin you want to search")
		} else {
			url := "http://vimawesome.com/api/plugins?query=" + os.Args[2]
			json, jsonError := GetJson(url)

			if nil != jsonError {
				fmt.Println("error")
			}

			response, parseError := ParsePluginsList(json)

			if nil != parseError {
				fmt.Println("error")
			}

			for _, plugin := range response.Plugins {
				fmt.Println("* " + plugin.Slug + " - " + plugin.ShortDesc)
			}
		}
	},
}
