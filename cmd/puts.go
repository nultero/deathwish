package cmd

import (
	"novem/cmd/argpkgs"
	"novem/cmd/fsys"
	"novem/cmd/index"
	"novem/cmd/put"
	"os"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var putsCmd = &cobra.Command{
	Use:     "puts [FILES]",
	Short:   argpkgs.PutsDesc,
	Aliases: argpkgs.PutsAliases,
	Args:    cobra.MinimumNArgs(1),
	Run:     puts,
}

// TODOOOO puts should be able to throw an error without crashing
// [x] empty index does not crash novem, only throws warning

func puts(cmd *cobra.Command, args []string) {

	dir, ihndl := viper.GetString(dataDir), viper.GetString(idxFile)
	ihndl = fsys.MeshPaths(dir, ihndl)
	idx := index.Init(ihndl)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		tics.ThrowSys(puts, err)
	}

	for _, arg := range args {
		put.Cmd(dir, arg, homeDir, RecurseFlag, &idx)
	}
}

func init() {
	rootCmd.AddCommand(putsCmd)
}
