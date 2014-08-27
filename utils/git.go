package utils

import "os/exec"

func GitClone(url string, path string, branch string) error {
	_, err := exec.Command("git", "clone", url, path).Output()

	if err != nil {
		return err
	}

	return nil
}
