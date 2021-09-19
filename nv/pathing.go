package nv

import (
	"fmt"
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

		if err != nil { //\\\\\\
			fmt.Println("exiting with error:", err)
			os.Exit(1)
		}

		path = p
	}

	return path
}
