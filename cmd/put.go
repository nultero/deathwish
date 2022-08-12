package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "add",
	Short: "adds dotfiles to the Deathwish registry | \x1b[34mrecursive\x1b[0m by default",
	Aliases: []string{
		"put",
		"puts",
	},
	Run: puts,
}

func init() {
	rootCmd.AddCommand(putCmd)
}

func puts(cmd *cobra.Command, args []string) {
	for _, arg := range args {
		err := uvarStat.Dotfiles.Add(arg, uvarStat.Home)
		if err != nil {
			fmt.Println(err)
		}
	}

	uvarStat.Close()
}
