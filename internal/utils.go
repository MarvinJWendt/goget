package internal

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/MarvinJWendt/goget/modules"
	"github.com/pterm/pterm"
	"os/exec"
	"strings"
)

func ModulesToDropdown() (entries *survey.Select) {
	entries = &survey.Select{
		Message: "Select a package to install:",
	}
	for _, pkg := range modules.Modules {
		var categories string
		for i, category := range pkg.Tags {
			if i == 0 {
				categories = string(category)
			} else {
				categories = categories + ", " + string(category)
			}
		}
		entries.Options = append(entries.Options, pterm.Sprintf("%s | %s | %s", pkg.Name, categories, pkg.Path))
	}

	return
}

func FilterVersionsToMajor(versions []string) (filtered []string) {
	for _, version := range versions {
		parts := strings.Split(version, ".")
		major := parts[0]
		if strings.HasPrefix(major, "v") {
			major = major[1:]
		}
		filtered = append(filtered, major)
	}
	return
}

func GetModuleByName(pkgName string) (pkg modules.Module, err error) {
	pkgName = strings.Split(pkgName, " ")[0]
	for _, pkg := range modules.Modules {
		if pkg.Name == pkgName {
			return pkg, nil
		}
	}
	return modules.Module{}, UNKNOWN_PACKAGE
}

func ExecGoGet(pkgPath string) (output string, err error) {
	bytes, err := exec.Command("go", "get", "-u", pkgPath).CombinedOutput()
	output = string(bytes)
	return
}

func InstallModule(pkgPath string) (err error) {
	pterm.Info.Println("Installing package:", pkgPath)
	output, err := ExecGoGet(pkgPath)
	if err != nil {
		return fmt.Errorf("error executing go get: %s (%w)", output, err)
	}

	return nil
}
