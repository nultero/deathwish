package cmd

import (
	"fmt"

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
	strs := uvarStat.Dotfiles.GetList()
	for _, s := range strs {
		fmt.Println(s)
	}
}
