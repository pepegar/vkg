package commands

import "testing"

func TestHasName(t *testing.T) {
	c := Command{
		Name:        "asdf",
		Description: "asdf",
		Usage:       "asdf",
		Action: func() {
		},
	}

	if !c.HasName("asdf") {
		t.Errorf("expected command to have asdf name, %s found", c.Name)
	}
}
