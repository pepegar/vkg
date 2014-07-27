package commands

type Command struct {
	Name        string
	Description string
	Action      func()
}

func (c *Command) HasName(name string) bool {
	return c.Name == name
}
