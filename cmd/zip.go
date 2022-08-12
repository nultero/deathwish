package cmd

import (
	"archive/zip"
	"deathwish/pathlib"
	"deathwish/uvars"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var zipOutFile = ""

var zipCmd = &cobra.Command{
	Use:   "zip",
	Short: "A brief description of your command",
	Run:   zipdots,
}

func init() {
	rootCmd.AddCommand(zipCmd)
	zipCmd.Flags().StringVarP(&zipOutFile, "outfile", "o", "", "the output zip archive file")
}

func zipdots(cmd *cobra.Command, args []string) {
	cwd, _ := os.Getwd()
	doChecks()
	os.Chdir(uvarStat.Home)
	file, err := os.Create(zipOutFile)
	if err != nil {
		panic(err)
	}
	zw := zip.NewWriter(file)
	defer zw.Close()

	dots := uvarStat.Dotfiles.GetCanonList(uvarStat.Home)
	zipDotList(dots, zw, uvarStat.Home)

	os.Chdir(cwd)
}

// Does all the ops on the dotfiles to zip them.
// Lots of repetitive err checking.
//
// No surprises here.
func zipDotList(dots []string, zw *zip.Writer, home string) {

	for _, f := range dots {

		finfo, err := os.Stat(f)
		if err != nil {
			fmt.Println(uvars.ERRHEADER, err)
			continue
		}

		hdr, err := zip.FileInfoHeader(finfo)
		if err != nil {
			fmt.Println(uvars.ERRHEADER, err)
			continue
		}
		hdr.Method = zip.Deflate

		idx := pathlib.GetIndexFmt(f, home)
		hdr.Name = idx

		hw, err := zw.CreateHeader(hdr)
		if err != nil {
			fmt.Println(uvars.ERRHEADER, err)
			continue
		}

		fp, err := os.Open(f)
		if err != nil {
			fmt.Println(uvars.ERRHEADER, err)
			continue
		}
		defer fp.Close()

		_, err = io.Copy(hw, fp)
		if err != nil {
			fmt.Println(uvars.ERRHEADER, err)
			continue
		}

	}
}

func doChecks() {
	if len(zipOutFile) == 0 {
		fmt.Println("\x1b[34m No outfile specified \x1b[0m")
		fmt.Println("please specify one with \x1b[34m-o [filename]\x1b[0m")
		return
	}
	zipAddExt()
}

func zipAddExt() {
	if zipOutFile[len(zipOutFile)-4:] != ".zip" {
		zipOutFile += ".zip"
	}
}
