package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	TotalResults   int `json:"total_results"`
	ResultsPerPage int `json:"results_per_page"`
	TotalPages     int `json:"total_pages"`
	Plugins        []PluginRecord
}

type PluginRecord struct {
	Author    string
	Slug      string
	ShortDesc string `json:"short_desc"`
	GithubUrl string `json:"github_url"`
}

func GetJson(url string) ([]byte, error) {
	resp, requestError := http.Get(url)

	if requestError != nil {
		return nil, requestError
	}

	defer resp.Body.Close()
	body, readError := ioutil.ReadAll(resp.Body)

	if readError != nil {
		return nil, readError
	}

	return body, nil
}

func ParsePluginsList(body []byte) (Response, error) {
	var r Response

	parseJsonError := json.Unmarshal(body, &r)

	return r, parseJsonError
}

func ParseSinglePlugin(body []byte) (PluginRecord, error) {
	var p PluginRecord

	parseJsonError := json.Unmarshal(body, &p)

	return p, parseJsonError
}
