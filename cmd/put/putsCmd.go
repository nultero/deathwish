package put

import (
	"fmt"
	"io/fs"
	"novem/cmd/fsys"
	"novem/cmd/index"
	"os"
	"path/filepath"

	"github.com/nultero/tics"
)

const emptStr = ""

func Cmd(nvDir, fpath, homeDir string, recurseFlag bool, idx *index.Index) {

	s := tics.Make(fpath)
	abs, finfo, err := chkPath(fpath)

	if idx.HasFile(abs) {
		hasNode, err := idx.HasInode(finfo)
		if err != nil {
			tics.ThrowSys(Cmd, err)
		}

		if hasNode {
			str := tics.Make("> novem already has file: ").Yellow()
			fmt.Printf("%v%v\n", str, abs)
			return
		} else if !hasNode {
			/*
				A file passed in as arg
				and a novem file have the
				same path, but their inodes
				do not match.

				Won't try to hardlink over,
				but print anyway.
			*/

			str := fmt.Sprintf("%v", tics.Make(abs).Yellow())
			fmt.Println(
				"filepath: '" + str + "' already" +
					"in novem's index, but is not" +
					" the same file as the argument",
			)
			return
		}
	}

	if finfo.IsDir() && !recurseFlag {
		fmt.Printf(
			"? %v: is directory, but '-r' flag is not set\n",
			s.Pink(),
		)
		return
	}

	tr, base := fsys.Trail(abs, homeDir)
	tr = fsys.AppendSlash(tr)

	dest := nvDir + tr + base

	err = fsys.NovemHardLink(abs, base, nvDir, dest)

	if err == nil {
		fmt.Printf("+ %v\n", s.Green())
		// TODOOOOOOOOO update index here

	} else {
		fmt.Printf("+ %v: %v\n", s.Red(), err)
	}
}

func chkPath(fpath string) (string, fs.FileInfo, error) {
	abs, err := filepath.Abs(fpath)
	if err != nil {
		return emptStr, nil, err
	}

	fh, err := os.Open(abs)
	if err != nil {
		return emptStr, nil, err
	}
	defer fh.Close()

	f, err := fh.Stat()
	if err != nil {
		return emptStr, nil, err
	}

	return abs, f, nil
}
