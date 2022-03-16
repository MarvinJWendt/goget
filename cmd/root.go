package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/MarvinJWendt/goget/internal"
	"os"
	"os/signal"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goget [modules]",
	Short: "A simple go module downloader",
	Long: `Goget is a simple go module downloader. If ran without arguments, it will prompt for a list of modules to get. If ran with arguments, it will get the specified modules.

If you want to add a missing module, you can do so by creating a PR on the goget repository: https://github.com/MarvinJWendt/goget`,
	Example: `goget
goget pterm
goget testza
goget pterm testza`,
	Version: "v0.1.0", // <---VERSION---> Updating this version, will also create a new GitHub release.
	Args:    cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			var pkgSelection string
			err := survey.AskOne(internal.ModulesToDropdown(), &pkgSelection)
			if err != nil {
				return err
			}
			pterm.Debug.Println("You selected:", pkgSelection)

			pkg := internal.GetModuleByName(pkgSelection)
			pterm.Debug.Printfln("Module: %#v", pkg)

			err = internal.InstallModule(pkg.Path)
			if err != nil {
				return err
			}
		} else {
			for _, arg := range args {
				pkg := internal.GetModuleByName(arg)

				pterm.Debug.Printfln("Module: %#v", pkg)

				err := internal.InstallModule(pkg.Path)
				if err != nil {
					return err
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
		pcli.CheckForUpdates()
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		pcli.CheckForUpdates()
		os.Exit(1)
	}

	pcli.CheckForUpdates()
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empty strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "d", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
	rootCmd.PersistentFlags().BoolVarP(&pcli.DisableUpdateChecking, "disable-update-checks", "", false, "disables update checks")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRepo("MarvinJWendt/goget")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
