package putsFn

import (
	"fmt"
	"novem/cmd/fsys"
	"os"
	"path/filepath"
	"strings"

	"github.com/nultero/tics"
)

// TODOOO update index
// TODOOO handle directory paths

const emptStr = ""

func Cmd(dir, fpath, homeDir string, recurse bool) {

	abs, base, name, blockRecur := checkPath(fpath, homeDir, recurse)
	if blockRecur {
		s := tics.Make(fpath).Pink().String()
		fmt.Printf("? %v: is directory, but '-r' flag is not set\n", s)
		return
	}

	dest := dir + base + name

	err := fsys.NovemHardLink(abs, base, dir, dest)
	fpath = filepath.Base(fpath) // slightly nicer to print

	if err == nil {
		s := tics.Make(fpath).Green().String()
		fmt.Printf("+ %v\n", s)

	} else {
		s := tics.Make(fpath).Red().String()
		fmt.Printf("+ %v: %v\n", s, err)
	}
}

// Unique to puts. Turn path into absolute filepath, then
// create a base novem variant of the path (with $HOME
//	stripped out) to hardlink.
func checkPath(path, home string, recurse bool) (string, string, string, bool) {
	abs, err := filepath.Abs(path)
	if err != nil {
		tics.ThrowSys(checkPath, err)
	}

	f, err := os.Open(abs)
	if err != nil {
		tics.ThrowSys(checkPath, err)
	}
	defer f.Close()

	st, err := f.Stat()
	if err != nil {
		tics.ThrowSys(checkPath, err)
	}

	if st.IsDir() && !recurse {
		return emptStr, emptStr, emptStr, true
	}

	base, name := filepath.Dir(abs), filepath.Base(abs)
	base = strings.ReplaceAll(base, home, emptStr)
	if base[0] == '/' {
		base = base[1:]
	}

	return abs, base, name, false
}
