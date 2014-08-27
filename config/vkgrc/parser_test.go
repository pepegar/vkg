package vkgrc

import "testing"

func TestParseVkgrc(t *testing.T) {
	fakeVkgrcContents := []byte(`{
		"plugins": [
			{
				"repository": "https://github.com/asdf/qwer",
				"name": "asdf",
				"branch": "master"
			}
		]
	}`)

	vkgrc := ParseVkgrc(fakeVkgrcContents)

	if len(vkgrc.Plugins) < 1 {
		t.Error("ParseVkgrc expects fakeVkgrcContents to be a valid vkgrc but isn't")
	}
}
