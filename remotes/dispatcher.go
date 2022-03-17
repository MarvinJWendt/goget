package remotes

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/MarvinJWendt/goget/internal"
	"github.com/pterm/pterm"
	"strconv"
	"strings"
)

func Run(remote string, arg string) error {
	if remote == "github" {
		repo, err := getRemote(arg)
		if err != nil {
			return err
		}
		pkgSelection, err := dropdown(repo)
		if err != nil {
			return fmt.Errorf("error selecting package  %w", err)
		}
		pkgPath := strings.Split(strings.Split(pkgSelection, " ")[2], ":")[1][2:]
		version, err := getVersion(strings.Split(pkgSelection, " ")[0])
		if err != nil {
			return err
		}
		if version[0:1] == "v" {
			i, err := strconv.Atoi(version[1:2])
			if err == nil {
				if i > 1 {
					pkgPath += "/" + version
					pterm.Debug.Println("Version tag higher than one")
				}
			}
		}
		pterm.Debug.Println(pkgPath)
		err = internal.InstallModule(pkgPath)
		return err
	} else {
		return errors.New("remote not implemented")
	}
}

func dropdown(repos *[]Repo) (string, error) {
	var pkgSelection string
	entries := &survey.Select{
		Message: "Select a package to install:",
	}
	for _, repo := range *repos {
		entries.Options = append(entries.Options, pterm.Sprintf("%s | %s | %s", repo.FullName, repo.HtmlUrl, repo.Description))
	}
	err := survey.AskOne(entries, &pkgSelection)
	if err != nil {
		return "", err
	}
	return pkgSelection, err
}
