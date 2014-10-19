package utils

import (
	"os"
	"os/exec"
	"strings"

	"github.com/pepegar/vkg/config"
)

type git struct{}

func (g git) Clone(url string, branch string) error {
	vkgConfig := config.GetVkgGonfig()
	os.Chdir(vkgConfig.PluginsPath)
	_, err := exec.Command("git", "clone", url).Output()

	if err != nil {
		return err
	}

	return nil
}

func (g git) GetBranchName(path string) (string, error) {
	os.Chdir(path)
	response, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	return trimLineJumps(string(response)), err
}

func (g git) GetRepository(path string) (string, error) {
	os.Chdir(path)
	response, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	return trimLineJumps(string(response)), err
}

func trimLineJumps(val string) string {
	return strings.Trim(val, "\n")
}

var Git = git{}
