package fsys

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"syscall"

	"github.com/nultero/tics"
)

func StatDir(dirName string, subdir *bool) ([]string, []int) {

	dir, err := ioutil.ReadDir(dirName)
	if err != nil {
		tics.ThrowSys(StatDir, err)
	}

	files := []string{}
	inodes := []int{}

	for _, f := range dir {
		s := f.Sys()
		if stat, ok := s.(*syscall.Stat_t); ok {
			files = append(files, f.Name())
			inodes = append(inodes, int(stat.Ino))
		}

		if *subdir {
			if f.IsDir() {
				subName := dirName + "/" + f.Name()
				subfiles, subnodes := StatDir(subName, subdir)
				files = append(files, subfiles...)
				inodes = append(inodes, subnodes...)
			}
		}
	}

	return files, inodes
}

func Inodes(args []string) ([]string, []int) {
	files := []string{}
	inodes := []int{}

	for _, arg := range args {
		file, err := filepath.Abs(arg)
		if err != nil {
			tics.ThrowUnhandled(fmt.Errorf("file `%v` can't be expanded into absolute path", file))
		}

		files = append(files, file)
	}

	return files, inodes
}
