package utils

import (
	"os"
	"os/exec"
	"strings"
)

type git struct{}

func (g git) Clone(url string, path string, branch string) error {
	_, err := exec.Command("git", "clone", url, path).Output()

	if err != nil {
		return err
	}

	return nil
}

func (g git) GetBranchName(path string) (string, error) {
	cwd, _ := os.Getwd()
	os.Chdir(path)
	response, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	os.Chdir(cwd)
	return trimLineJumps(string(response)), err
}

func (g git) GetRepository(path string) (string, error) {
	cwd, _ := os.Getwd()
	os.Chdir(path)
	response, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	os.Chdir(cwd)
	return trimLineJumps(string(response)), err
}

func (g git) GetRepoName(path string) (string, error) {
	cwd, _ := os.Getwd()
	os.Chdir(path)
	response, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	os.Chdir(cwd)
	return trimLineJumps(basename(string(response))), err
}

func trimLineJumps(val string) string {
	return strings.Trim(val, "\n")
}

func basename(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

var Git = git{}
