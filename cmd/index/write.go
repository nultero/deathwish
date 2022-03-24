package index

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/nultero/tics"
)

func (idx *Index) Write(idxpath string) {

	strs := []string{}
	for k, i := range idx.PathMatches {
		no := fmt.Sprint(idx.Entries[i].Inode)
		strs = append(strs, fmt.Sprintf(
			"%v%v%v%v%v%v%v",
			k, charSep,
			no, charSep,
			idx.Entries[i].OutLinkPath, charSep,
			idx.Entries[i].ChangedLast,
		))
	}

	sort.Strings(strs)
	out := strings.Join(strs, "\n")
	err := os.WriteFile(idxpath, []byte(out), tics.Perm)
	if err != nil {
		s := tics.Make("hardlinks performed but no information saved, index write failed: ").Red().String()
		err = fmt.Errorf("%v%w", s, err)
		tics.ThrowSys(idx.Write, err)
	}
}
