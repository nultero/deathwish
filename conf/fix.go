package conf

import (
	"errors"
	"fmt"
	"io/fs"
	"novem/colors"
	"novem/nv"
	"os"
)

func Fix(path string) {

	_, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			novemDirNotExists(path)
		}
	}

}

func novemDirNotExists(path string) {
	dir := nv.GetHomeDir() + nv.ExpandTilde(path)
	s := "'" + colors.Blue(dir) + "'"
	fmt.Println("novem directory", s, "does not exist")
}
