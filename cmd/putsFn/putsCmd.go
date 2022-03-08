package putsFn

import (
	"fmt"
	"novem/cmd/fsys"
	"os"
	"path/filepath"
	"strings"

	"github.com/nultero/tics"
)

// TODOOOO check index +
// TODOOO update index
// TODOOO handle directory paths

const emptStr = ""

func Cmd(nvDir, fpath, homeDir string, recurse bool) {

	//check index first

	s := tics.Make(fpath)
	abs, isDir, err := chkPath(fpath)

	if isDir && !recurse {
		fmt.Printf(
			"? %v: is directory, but '-r' flag is not set\n",
			s.Pink().String(),
		)
		return
	}

	tr, base := fsys.Trail(abs, homeDir) // check index here?
	tr = fsys.AppendSlash(tr)

	dest := nvDir + tr + base // TODOOOOOO write some tests & funcs to handle the permutations here
	fmt.Println(dest, "+ dest")

	// err := fsys.NovemHardLink(abs, base, dir, dest)

	if err == nil {
		fmt.Printf("+ %v\n", s.Green().String())

	} else {
		fmt.Printf("+ %v: %v\n", s.Red().String(), err)
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

func pthTrail(abs, home string) (string, string) {
	tr, base := filepath.Dir(abs), filepath.Base(abs)
	tr = strings.ReplaceAll(tr, home, emptStr)
	if tr[0] == '/' {
		tr = tr[1:]
	}
	return tr, base
}
