package nv

import (
	"novem/errs"
	"os"
	"path/filepath"
)

func GetHomeDir() string {
	h, err := os.UserHomeDir()

	if err != nil {
		errs.SysErr(err)
	}

	return h
}

func ExpandTilde(path string) string {

	if path[0] == '~' {
		path = path[1:]
		p, err := filepath.Abs(path)

		if err != nil {
			errs.SysErr(err)
		}

		path = p
	}

	return path
}

func GetFullPath(path string) string {
	return GetHomeDir() + ExpandTilde(path)
}
