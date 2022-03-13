package fsys

import (
	"errors"
	"io/fs"
	"os"
	"strings"

	"github.com/nultero/tics"
)

func NovemHardLink(fileName, novemDir, dest string) error {

	strip := stripNovDir(dest, novemDir)
	tr := subTrail(strip)

	if parentDirExists(novemDir + tr) {
		err := link(fileName, dest)
		if err != nil {
			return err
		}
	} else {
		err := mkDirs(tr, novemDir)
		if err != nil {
			return err
		}

		err = link(fileName, dest)
		if err != nil {
			return err
		}
	}

	return nil
}

func link(fn, dest string) error {
	err := os.Link(fn, dest)
	if err != nil {
		return err
	}
	return nil
}

func parentDirExists(dir string) bool {
	f, err := os.Stat(dir)
	if err != nil {
		return false
	}

	if f.IsDir() {
		return true
	}

	return false
}

func stripNovDir(dest, novDir string) string {
	d := strings.Replace(dest, novDir, "", 1)
	if d[0] == '/' {
		d = d[1:]
	}
	return d
}

func mkDirs(trail, nvDir string) error {
	dirs := strings.Split(trail, "/")
	p := ""
	for _, d := range dirs {
		path := nvDir + p + d
		err := os.Mkdir(path, tics.Perm)
		if errors.Is(err, fs.ErrExist) {
			return nil
		} else if err != nil {
			return err
		}

		p += d + "/"
	}
	return nil
}
