package cmd

import (
	"archive/zip"
	"deathwish/pathlib"
	"deathwish/uvars"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var unzipCmd = &cobra.Command{
	Use:   "unzip",
	Short: "decompress all files in the given .deathzip file",
	Run:   unzipMain,
}

func unzipMain(cmd *cobra.Command, args []string) {
	uzArgChecks(args)
	rdr, err := zip.OpenReader(args[0])
	if err != nil {
		panic(err)
	}

	for _, zf := range rdr.File {
		uzFile(zf, uvarStat.Home)
	}
}

func uzFile(zf *zip.File, home string) {
	hdr := zf.FileHeader
	fpath := pathlib.SubCanon(hdr.Name, home)

	err := os.MkdirAll(filepath.Dir(fpath), 0755)
	if err != nil {
		fmt.Println(
			uvars.ERRHEADER,
			err,
		)
		return
	}

	f, err := os.Create(fpath)
	if err != nil {
		fmt.Println(
			uvars.ERRHEADER,
			err,
		)
		return
	}
	defer f.Close()

	rdr, err := zf.Open()
	if err != nil {
		fmt.Println(
			uvars.ERRHEADER,
			err,
		)
		return
	}
	defer rdr.Close()
	_, err = io.Copy(f, rdr)
	if err != nil {
		fmt.Println(
			uvars.ERRHEADER,
			err,
		)
		return
	}
}

func uzArgChecks(args []string) {
	quit := false
	if len(args) == 0 {
		fmt.Println(uvars.ERRHEADER, "not enough args")
		quit = true
	}

	if len(args) > 1 {
		fmt.Println(uvars.ERRHEADER, "multiple zips???")
		quit = true
	}

	if quit {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(unzipCmd)
}
