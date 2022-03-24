package index

import (
	"io/fs"
	"syscall"
)

func (idx *Index) Update(abs, nvDest string, finfo fs.FileInfo) {
	stat, _ := finfo.Sys().(*syscall.Stat_t)
	mt := finfo.ModTime().String()
	li := len(idx.Entries)

	e := Entry{
		NovemPath:   nvDest,
		Inode:       stat.Ino,
		OutLinkPath: abs,
		ChangedLast: mt,
	}

	idx.Entries = append(idx.Entries, e)
	idx.PathMatches[nvDest] = li
	idx.PathSet[abs] = struct{}{}
}
