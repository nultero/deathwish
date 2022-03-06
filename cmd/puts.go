package cmd

import (
	"novem/cmd/argpkgs"
	"novem/cmd/putsFn"
	"os"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"
)

var putsCmd = &cobra.Command{
	Use:     "puts [FILES]",
	Short:   argpkgs.PutsDesc,
	Aliases: argpkgs.PutsAliases,
	Args:    cobra.MinimumNArgs(1),
	Run:     puts,
}

// TODOOO puts should be able to throw an error without crashing

func puts(cmd *cobra.Command, args []string) {
	if dir, ok := confMap[dataDir]; ok {
		dir += "/"
		homeDir, err := os.UserHomeDir()
		if err != nil {
			tics.ThrowSys(puts, err)
		}

		for _, arg := range args {
			putsFn.Cmd(dir, arg, homeDir, RecurseFlag)
		}
	}
}

func init() {
	rootCmd.AddCommand(putsCmd)
}
