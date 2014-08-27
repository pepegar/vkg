package commands

import "testing"

func TestIsUserRepo(t *testing.T) {
	validUserRepos := []string{
		"asdf/qwer",
		"qwer-qwer/qwer",
		"asdf/wqer-qwer",
	}

	inValidUserRepos := []string{
		"asdf/",
		"/qwer",
		"-qwer",
	}

	for _, item := range validUserRepos {
		if !IsUserRepo(item) {
			t.Errorf("expected %s to be valid user repo but seems to be invalid", item)
		}
	}

	for _, item := range inValidUserRepos {
		if IsUserRepo(item) {
			t.Errorf("expected %s to be invalid user repo but seems to be valid", item)
		}
	}
}

func TestIsGithubUrl(t *testing.T) {
	validGithubUrls := []string{
		"github.com/asdf/qwer",
		"github.com/qwer-qwer/qwer",
		"github.com/asdf/wqer-qwer",
	}

	invalidGithubUrls := []string{
		"github.com/asdf/",
		"github.com//qwer",
		"github.com/-qwer",
	}

	for _, item := range validGithubUrls {
		if !IsGithubUrl(item) {
			t.Errorf("expected %s to be valid github url but seems to be invalid", item)
		}
	}

	for _, item := range invalidGithubUrls {
		if IsGithubUrl(item) {
			t.Errorf("expected %s to be invalid github url but seems to be valid", item)
		}
	}
}
