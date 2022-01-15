package cmd

import (
	"fmt"
	"novem/cmd/fsys"
	"path/filepath"

	"github.com/spf13/cobra"
)

// putsCmd represents the puts command
var putsCmd = &cobra.Command{
	Use:   "puts [FILES]",
	Short: "A brief description of your command",
	Args:  cobra.MinimumNArgs(1),
	Run:   puts,
}

func puts(cmd *cobra.Command, args []string) {
	if dir, ok := confMap[dataDir]; ok {
		dir += "/"
		for _, arg := range args {
			abs, base := checkPath(arg)
			dest := dir + base
			err := fsys.NvHardLink(abs, dest)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func checkPath(p string) (string, string) {
	s, err := filepath.Abs(p)
	if err != nil {
		return "", p
	}

	return s, filepath.Base(s)
}

func init() {
	rootCmd.AddCommand(putsCmd)
}
