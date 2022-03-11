package index

import (
	"errors"
	"fmt"
	"io/fs"
	"syscall"

	"github.com/nultero/tics"
)

func (idx *Index) HasFile(abs string) bool {
	_, has := idx.PathSet[abs]
	return has
}

// Tries to stat the file to find an inode, and so can trip an error if the stat fails.
func (idx *Index) HasInode(f fs.FileInfo) (bool, error) {

	s, ok := f.Sys().(*syscall.Stat_t)
	if !ok {
		return false, inoErr(f.Name())
	}

	inode := s.Ino

	for _, e := range idx.Entries {
		if inode == e.Inode {
			return true, nil
		}
	}

	return false, nil
}

func inoErr(fpath string) error {
	y := tics.Make(fpath).Yellow()
	s := fmt.Sprintf("syscall stat failed on file: %v", y)
	return errors.New(s)
}
