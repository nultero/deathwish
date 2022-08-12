package cmd

import (
	"deathwish/uvars"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var uvarStat = uvars.Uvars_t{}

var rootCmd = &cobra.Command{
	Use:   "deathwish",
	Short: "Some brief desc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(uvarStat)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	u, err := uvars.InitUvars()
	if err != nil {
		panic(err)
	}

	uvarStat = u
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
