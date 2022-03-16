package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/MarvinJWendt/goget/modules"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Prints all inbuild modules as json",
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := json.Marshal(modules.Modules)
		if err != nil {
			return err
		}

		fmt.Println(string(b))
		return err
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
