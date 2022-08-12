package uvars

import (
	"deathwish/pathlib"
	"fmt"
	"os"
	"sort"
)

type Dots_t struct {
	Static  map[string]struct{}
	NewDots []string
}

func initDots() Dots_t {
	return Dots_t{
		Static:  map[string]struct{}{},
		NewDots: []string{},
	}
}

// This also fails for edge cases like
// whether user has permissions to read X
// file.
func fileExists(path string) bool {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return false
	}
	return true
}

func (d Dots_t) GetList() []string {
	i := 0
	for f := range d.Static {
		if len(f) != 0 {
			i++
		}
	}

	files := make([]string, i)
	i = 0
	for f := range d.Static {
		files[i] = f
		i++
	}

	sort.Strings(files)
	return files
}

func (d Dots_t) GetCanonList(home string) []string {
	ls := d.GetList()
	for i, f := range ls {
		ls[i] = pathlib.SubCanon(f, home)
	}
	return ls
}

func (d *Dots_t) Add(fpath, home string) error {
	cpath, err := pathlib.GetCanon(fpath)
	if err != nil {
		return err
	}

	if !fileExists(cpath) {
		err = fmt.Errorf(
			"%s file \x1b[;1m%s\x1b[0m does not exist or deathwish cannot access it;\n%w",
			ERRHEADER,
			cpath, err,
		)
		return err
	}

	finfo, err := os.Stat(cpath)
	if err != nil {
		err = fmt.Errorf(
			"%s os.Stat() failed on file \x1b[;1m%s\x1b[0m;\n%w",
			ERRHEADER,
			cpath, err,
		)
		return err
	}

	if finfo.IsDir() {
		dir, err := os.ReadDir(cpath)
		if err != nil {
			err = fmt.Errorf(
				"%s read of directory failed on \x1b[;1m%s\x1b[0m;\n%w",
				ERRHEADER,
				cpath, err,
			)
		}

		for _, f := range dir {
			newPath := cpath + "/" + f.Name()
			err := d.Add(newPath, home)
			if err != nil {
				err = fmt.Errorf(
					"%s sub-error in recursive call on \x1b[;1m%s\x1b[0m;\n\tsub-error -> %w",
					ERRHEADER,
					f.Name(), err,
				)
				fmt.Println(err)
			}
		}

		return nil
	}

	uc := pathlib.GetIndexFmt(cpath, home)
	d.NewDots = append(d.NewDots, uc)
	return nil
}