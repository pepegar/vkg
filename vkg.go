package main

import (
	"fmt"
	"os"
)

type Command struct {
	Name        string
	Description string
	Action      func()
}

func main() {
	commands := []Command{
		Command{
			Name:        "test",
			Description: "test",
			Action: func() {
				println("asdf")
			},
		},
		Command{
			Name:        "test2",
			Description: "test2",
			Action: func() {
				println("asdf2")
			},
		},
	}

	if len(os.Args) > 1 {
		fmt.Println(commands)
	} else {
		fmt.Println("please, supply a valid command")
	}

}
