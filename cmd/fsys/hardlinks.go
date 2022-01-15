package fsys

import (
	"os"

	"github.com/nultero/tics"
)

func NvHardLink(filename, dest string) error {
	err := copy(filename, dest)
	if err != nil {
		return err
	}

	err = os.Remove(filename)
	if err != nil {
		return err
	}

	err = os.Link(dest, filename)
	if err != nil {
		return err
	}

	return nil
}

// Move filestate to novem for gittable repo.
func copy(fname, dest string) error {
	bytes, err := os.ReadFile(fname)
	if err != nil {
		return err
	}

	err = os.WriteFile(dest, bytes, tics.Perm)
	if err != nil {
		return err
	}

	return nil
}
