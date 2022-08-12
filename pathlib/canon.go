package pathlib

import (
	"path/filepath"
	"strings"
)

const tilde = "~"

func SubCanon(fpath, home string) string {
	return strings.Replace(fpath, tilde, home, 1)
}

func GetCanon(fpath string) (string, error) {
	f, err := filepath.Abs(fpath)
	if err != nil {
		return fpath, err
	}

	return f, nil
}

func GetIndexFmt(fpath, home string) string {
	return strings.Replace(fpath, home, tilde, 1)
}
