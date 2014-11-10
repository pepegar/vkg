package commands

import (
	"testing"

	"github.com/pepegar/vkg/config"
)

func TestHasName(t *testing.T) {
	c := Command{
		Name:        "asdf",
		Description: "asdf",
		Usage:       "asdf",
		Action: func(vkgConfig config.Config) {
		},
	}

	if !c.HasName("asdf") {
		t.Errorf("expected command to have asdf name, %s found", c.Name)
	}
}
