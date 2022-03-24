package index

import (
	"fmt"
	"path/filepath"

	"github.com/nultero/tics"
)

func (idx *Index) SimpleList(f string) error {

	full, err := filepath.Abs(f)
	if err != nil {
		return err
	}

	b := filepath.Base(f)
	f = "𝤎  " + b // TODOOO put the symbol thing into config

	if idx.HasFile(full) {
		// TODOOOO pass in verbosity level: last changed details, etc.
		if len(f) != 0 {
			fmt.Println(tics.Make(f).Cyan().String())
		}

	} else {
		// print file name grayed out?
		// 𐄂
	}

	return nil
}
