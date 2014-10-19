package vkgrc

import "encoding/json"

type VkgrcJSON struct {
	Plugins []VkgrcPlugin `json:"plugins"`
}

type VkgrcPlugin struct {
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
}

func ParseVkgrc(contents []byte) VkgrcJSON {
	var vkgrc VkgrcJSON

	json.Unmarshal(contents, &vkgrc)

	return vkgrc
}
