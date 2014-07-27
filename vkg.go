package main

import (
	"fmt"
	"os"

	"github.com/pepegar/vkg-go/commands"
)

func main() {
	commands := []commands.Command{
		commands.SearchCommand,
		commands.InstallCommand,
	}

	if len(os.Args) > 1 {
		fmt.Println(commands)
	} else {
		fmt.Println("please, supply a valid command")
	}

}
