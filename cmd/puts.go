package cmd

import (
	"fmt"
	"novem/cmd/argpkgs"
	"novem/cmd/fsys"
	"novem/cmd/index"

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

// TODOOO puts should be able to throw an error without crashing

func puts(cmd *cobra.Command, args []string) {

	dir, idxHndl := viper.GetString(dataDir), viper.GetString(idxFile)
	idxHndl = fsys.MeshPaths(dir, idxHndl)

	idx := index.Init(idxHndl)
	fmt.Println(idx) // TODOOOOOO pass in idx as dep inj to tests and puts.Cmd

	// if dir, ok := confMap[dataDir]; ok {
	// 	dir = fsys.AppendSlash(dir)
	// 	homeDir, err := os.UserHomeDir()
	// 	idx := index.From()

	// 	if err != nil {
	// 		tics.ThrowSys(puts, err)
	// 	}

	// 	for _, arg := range args {
	// 		putsFn.Cmd(dir, arg, homeDir, RecurseFlag)
	// 	}
	// }
}

func init() {
	rootCmd.AddCommand(putsCmd)
}
