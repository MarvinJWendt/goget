package remotes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MarvinJWendt/goget/internal"
	"github.com/pterm/pterm"
	"io/ioutil"
	"net/http"
	"strconv"
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

type Tag struct {
	Name   string `json:"name,omitempty"`
	NodeId string `json:"node_id,omitempty"`
}

func getRemote(arg string) (*[]Repo, error) {
	resp, err := http.Get("https://api.github.com/search/repositories?q=" + arg + "+language:go")
	if err != nil {
		return nil, fmt.Errorf("error getting response  %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("github did not respond with 200")
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
	if respDecoded.TotalCount == 0 {
		return nil, internal.UNKNOWN_PACKAGE
	}
	return &respDecoded.Items, err
}

func getVersion(repoName string) (string, error) {
	resp, err := http.Get("https://api.github.com/repos/" + repoName + "/tags")
	if err != nil {
		return "", fmt.Errorf("error getting response  %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return "", internal.UNKNOWN_PACKAGE
		}
		return "", errors.New("github did not respond with 200")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response  %w", err)
	}
	var tags []Tag
	err = json.Unmarshal(body, &tags)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response  %w", err)
	}
	for _, tag := range tags {
		if tag.Name[0:1] == "v" {
			i, err := strconv.Atoi(tag.Name[1:2])
			if err != nil {
				continue
			}
			pterm.Debug.Printfln("Found tag with version number: %s", tag)
			if i > 1 {
				return tag.Name[0:2], err
			}
		} else {
			continue
		}
	}
	pterm.Debug.Println("No tag with version number")
	return "", nil
}
