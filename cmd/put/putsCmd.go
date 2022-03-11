package put

import (
	"fmt"
	"novem/cmd/fsys"
	"novem/cmd/index"
	"os"
	"path/filepath"

	"github.com/nultero/tics"
)

// TODOOOO check index +
// TODOOO update index
// TODOOO handle directory paths

const emptStr = ""

func Cmd(nvDir, fpath, homeDir string, recurseFlag bool, idx *index.Index) {

	s := tics.Make(fpath)
	abs, isDir, err := chkPath(fpath)

	if idx.HasFile(abs) {
		// want to check if the nodes match if paths match
		// if nodes match, print "novem already has this file"
		// if not, then explain
	}
	//check index first

	if isDir && !recurseFlag {
		fmt.Printf(
			"? %v: is directory, but '-r' flag is not set\n",
			s.Pink(),
		)
		return
	}

	tr, base := fsys.Trail(abs, homeDir) // check index here?
	tr = fsys.AppendSlash(tr)

	dest := nvDir + tr + base // TODO flesh out the tests that handle the permutations here
	fmt.Println(dest, "+ dest")

	// err := fsys.NovemHardLink(abs, base, dir, dest)

	if err == nil {
		fmt.Printf("+ %v\n", s.Green())

	} else {
		fmt.Printf("+ %v: %v\n", s.Red(), err)
	}
}

func chkPath(fpath string) (string, bool, error) {
	abs, err := filepath.Abs(fpath)
	if err != nil {
		return emptStr, false, err
	}

	fh, err := os.Open(abs)
	if err != nil {
		return emptStr, false, err
	}
	defer fh.Close()

	f, err := fh.Stat()
	if err != nil {
		return emptStr, false, err
	}

	return abs, f.IsDir(), nil
}
