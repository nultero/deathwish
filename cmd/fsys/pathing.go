package fsys

import (
	"path/filepath"
	"strings"
)

const emptStr = ""

// Takes an absolute filepath & home path, breaks the absolute path
// into a directory trail and its file, and shaves the $home off of the
// trailpath, returns trail and base filepath.
//
// Strips extra starting '/' byte from the trailpath if necessary.
func Trail(abs, home string) (string, string) {
	tr, base := filepath.Dir(abs), filepath.Base(abs)
	tr = strings.ReplaceAll(tr, home, emptStr)
	if tr[0] == '/' {
		tr = tr[1:]
	}
	return tr, base
}

func subTrail(abs string) string {
	tr := filepath.Dir(abs)
	if tr[0] == '/' {
		tr = tr[1:]
	}
	return tr
}

// Takes a path and appends a '/' dirslash if it doesn't have one.
func AppendSlash(fpath string) string {
	if fpath[len(fpath)-1] != '/' {
		fpath += "/"
	}
	return fpath
}

func MeshPaths(dir, base string) string {
	c := 0
	if dir[len(dir)-1] == '/' {
		c++
	}

	if base[0] == '/' {
		c++
	}

	switch c {
	case 0:
		dir += "/"
	case 2:
		for base[0] == '/' {
			base = base[1:]
		}
	}

	return dir + base
}
