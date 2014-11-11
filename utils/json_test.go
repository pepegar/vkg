package utils

import (
	"testing"
)

func TestParsePluginsList(t *testing.T) {
	validJson := []byte(`{
	"total_results": 273,
	"results_per_page": 20,
	"total_pages": 14,
	"plugins": [
		{
			"author": "Tim Pope",
			"slug": "fugitive-vim",
			"short_desc": "fugitive.vim: a Git wrapper so awesome, it should be illegal",
			"github_url": "https://github.com/tpope/vim-fugitive"
		}
	]
	}`)

	response, parseError := ParsePluginsList(validJson)

	if parseError != nil {
		t.Error("parsing a valid json throws an error")
	}

	if response.TotalResults != 273 {
		t.Error(response.TotalResults)
	}

	invalidJson := []byte(`{asdf`)

	_, parseErrorTwo := ParsePluginsList(invalidJson)
	if parseErrorTwo == nil {
		t.Error("parsing invalid json doesn't throw an error")
	}
}
