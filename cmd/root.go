package cmd

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/MarvinJWendt/goget/internal"
	"github.com/MarvinJWendt/goget/modules"
	"github.com/MarvinJWendt/goget/remotes"
	"os"
	"os/signal"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var remote string

var rootCmd = &cobra.Command{
	Use:   "goget [modules]",
	Short: "A simple go module downloader",
	Long: `Goget is a simple go module downloader. If ran without arguments, it will prompt for a list of modules to get. If ran with arguments, it will get the specified modules.

If you want to add a missing module, you can do so by creating a PR on the goget repository: https://github.com/MarvinJWendt/goget`,
	Example: `goget
goget pterm
goget testza
goget pterm testza`,
	Version: "v1.2.1", // <---VERSION---> Updating this version, will also create a new GitHub release.
	Args:    cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		modules.FetchNewModules()

		if len(args) == 0 {
			var pkgSelection string
			err := survey.AskOne(internal.ModulesToDropdown(), &pkgSelection)
			if err != nil {
				return err
			}
			pterm.Debug.Println("You selected:", pkgSelection)

			pkg, _ := internal.GetModuleByName(pkgSelection)
			pterm.Debug.Printfln("Module: %#v", pkg)

			err = internal.InstallModule(pkg.Path)
			if err != nil {
				return err
			}
		} else {
			for _, arg := range args {
				if remote != "none" {
					err := remotes.Run(remote, arg)
					if err != nil {
						if errors.Is(err, internal.UNKNOWN_PACKAGE) {
							pterm.Warning.Printfln("Unable to find package \"%s\" on %s", arg, remote)
						} else {
							return err
						}
					}
				} else {
					pkg, err := internal.GetModuleByName(arg)
					if err != nil {
						if errors.Is(err, internal.UNKNOWN_PACKAGE) {
							pterm.Warning.Printfln("Unable to find package \"%s\", try -r", arg)
						} else {
							return err
						}
					} else {
						pterm.Debug.Printfln("Module: %#v", pkg)

						err = internal.InstallModule(pkg.Path)
						if err != nil {
							return err
						}
					}
				}
			}
		}

		pterm.Success.Println("Done!")

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Fetch user interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Warning.Println("user interrupt")
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empty strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "d", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
	rootCmd.PersistentFlags().StringVarP(&remote, "remote", "r", "none", "Look for module on git")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRepo("MarvinJWendt/goget")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
