package cmd

import (
	"fmt"
	"novem/cmd/fsys"
	"novem/cmd/index"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",

	Run: lsFunc,
}

// TODO docs, match list on path args
func lsFunc(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		s := tics.Make(" -- ").Cyan().String()
		fmt.Println(s + "no directories given to list")
		return
	}

	nvDir, ihndl := viper.GetString(dataDir), viper.GetString(idxFile)
	ihndl = fsys.MeshPaths(nvDir, ihndl)
	idx := index.Init(ihndl)

	for _, arg := range args {
		err := idx.SimpleList(arg)
		if err != nil { // may not be fatal
			fp := fmt.Sprintf("* problem listing:")
			arg = tics.Make(arg).Red().String()
			fmt.Println(fp, arg)
		}
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().CountVarP(&Verbosity, "verbose", "v", "increases the detail that index outputs")
}
