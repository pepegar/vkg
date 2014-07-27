package commands

type Command struct {
	Name        string
	Description string
	Action      func()
}
