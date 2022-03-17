package remotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Resp struct {
	TotalCount        int    `json:"total_count,omitempty"`
	IncompleteResults bool   `json:"incomplete_results,omitempty"`
	Items             []Repo `json:"items,omitempty"`
}

type Repo struct {
	Id          int    `json:"id,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	HtmlUrl     string `json:"html_url,omitempty"`
	Description string `json:"description,omitempty"`
}

func getRemote(arg string) (*[]Repo, error) {
	resp, err := http.Get("https://api.github.com/search/repositories?q=" + arg + "+language:go")
	if err != nil {
		return nil, fmt.Errorf("error getting response  %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response  %w", err)
	}
	var respDecoded Resp
	err = json.Unmarshal(body, &respDecoded)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response  %w", err)
	}
	return &respDecoded.Items, err
}
