package cmd

import (
	"fmt"
	"github.com/MarvinJWendt/goget/modules"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
)

var internalCICmd = &cobra.Command{
	Use:    "internal-ci",
	Hidden: true,
	Short:  "Used for internal CI",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.OpenFile("./README.md", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return err
		}

		defer f.Close()

		text := `
# Registered Modules

> Goget currently has **%d** modules registered.  
> To add modules into the goget registry, make a PR on GitHub!  
> https://github.com/MarvinJWendt/goget/blob/main/modules/modules.go  

|Name|URL|Tags|
|---|---|---|
`

		text = fmt.Sprintf(text, len(modules.Modules))

		for _, m := range modules.Modules {
			text += pterm.Sprintfln("|%s|https://%s|%s|", m.Name, m.Path, m.Tags)
		}

		if _, err = f.WriteString(text); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(internalCICmd)
}
