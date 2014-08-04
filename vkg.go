package main

import (
	"fmt"
	"os"

	"github.com/pepegar/vkg-go/commands"
)

type App struct {
	Commands []commands.Command
}

func (a *App) Dispatch(name string) *commands.Command {
	for _, c := range a.Commands {
		if c.HasName(name) {
			return &c
		}
	}

	return nil
}

func main() {
	app := App{
		Commands: []commands.Command{
			commands.SearchCommand,
			commands.InstallCommand,
			commands.UninstallCommand,
		},
	}

	var usage = `

Usage: vkg [command]

Commands:
`
	for _, command := range app.Commands {
		usage += "  " + command.Usage + " - " + command.Description + "\r\n"
	}

	if len(os.Args) > 1 {
		command := app.Dispatch(os.Args[1])

		if command != nil {
			command.Action()
		} else {
			fmt.Println("command " + os.Args[1] + " does not exist")
		}
	} else {
		fmt.Println(usage)
	}

}
