package modules

import (
	"encoding/json"
	"github.com/pterm/pterm"
	"io/ioutil"
	"net/http"
)

// FetchNewModules fetches all registered modules from GitHub
func FetchNewModules() {
	pterm.Debug.Println("Fetching new modules")
	resp, err := http.Get("https://raw.githubusercontent.com/MarvinJWendt/goget/main/generated/modules.json")
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var modules []Module

	err = json.Unmarshal(body, &modules)
	if err != nil {
		return
	}

	Modules = modules
}
