package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",

	Run: lsFunc,
}

// TODO docs, match list on path args
func lsFunc(cmd *cobra.Command, args []string) {

	// if path, ok := confMap[dataDir]; ok {

	// 	idx, err := index.Get(path)
	// 	if err != nil {
	// 		if os.IsNotExist(err) {
	// 			index.WriteIndex(path, map[int]interface{}{})

	// 		} else {
	// 			tics.ThrowSys(lsFunc, err)
	// 		}
	// 	}

	// 	idx.List(Verbosity)
	// }
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().CountVarP(&Verbosity, "verbose", "v", "increases the detail that index outputs")
}
