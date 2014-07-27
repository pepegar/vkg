package main

import (
	"fmt"
	"os"

	"github.com/pepegar/vkg-go/commands"
)

type App struct {
	Commands []commands.Command
}

func (a *App) Command(name string) *commands.Command {
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
		},
	}

	if len(os.Args) > 1 {
		command := app.Command(os.Args[1])

		if command != nil {
			command.Action()
		} else {
			fmt.Println("command " + os.Args[1] + " does not exist")
		}
	} else {
		fmt.Println("please, supply a valid command")
	}

}
