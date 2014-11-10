package commands

import "github.com/pepegar/vkg/config"

type Command struct {
	Name        string
	Description string
	Usage       string
	Action      func(config.Config)
}

func (c *Command) HasName(name string) bool {
	return c.Name == name
}
