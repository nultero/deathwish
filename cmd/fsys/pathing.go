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
