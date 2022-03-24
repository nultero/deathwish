package index

import (
	"path/filepath"
)

func (idx *Index) SimpleList(f string) error {

	full, err := filepath.Abs(f)
	if err != nil {
		return err
	}

	if idx.HasFile(full) {
		// prints cyan
	} else {
		// grayed out
	}

	return nil
}
