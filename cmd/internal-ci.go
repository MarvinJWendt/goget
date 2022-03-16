package cmd

import (
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

|Name|URL|Tags|
|---|---|---|
`

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
