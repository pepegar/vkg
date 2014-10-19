package utils

import "os/exec"

type git struct{}

func (g git) Clone(url string, path string, branch string) error {
	_, err := exec.Command("git", "clone", url, path).Output()

	if err != nil {
		return err
	}

	return nil
}

var Git = git{}
