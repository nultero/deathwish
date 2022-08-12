package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "A brief description of your command",
	Run:     list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {
	strs := make([]string, len(uvarStat.Dotfiles.Static))
	i := 0
	for f := range uvarStat.Dotfiles.Static {
		strs[i] = f
		i++
	}

	sort.Strings(strs)

	for _, s := range strs {
		fmt.Println(s)
	}
}
