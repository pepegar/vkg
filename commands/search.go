package commands

var SearchCommand = Command{
	Name:        "search",
	Description: "Search a plugin",
	Action: func() {
		println("Search!")
	},
}
