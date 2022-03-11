package index

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nultero/tics"
)

// The index file is, more or less, the declarative state of
// what the dots should look like. Accidental moves and deletions
// should crop up here, and it should be a little bit faster than
// trying to read a dir/dirs into memory.
// Probably have to use hardlinks to be able to track certain
// things like file moves, so this does inherently limit novem
// to a single local machine. Could use git to distribute deltas
// across machines.
var IdxFile = "/.novem_index"

type Index struct {
	PathSet map[string]int // these are the novem dir's paths
	Entries []Entry
}

// Reads index into memory, allows operations on
// idx. Hard crash if this fails; I want this to
// be obvious when this goes wrong.
func Init(fpath string) Index {
	bytes, err := os.ReadFile(fpath)
	if err != nil {
		tics.ThrowSys(Init, err)
	}

	s := string(bytes)
	if len(s) == 0 {
		return uninitWarn(fpath)
	}

	// if the index is otherwise invalid, I do think I want a panic
	lines := strings.Split(s, "\n")
	idx := initEmptIdx()

	for i, ln := range lines {
		el := strings.Split(ln, charSep)
		if len(el) != 4 {
			tics.ThrowSys(Init, indexErr(fpath, i))
		}

		inode, err := strconv.ParseUint(el[1], 10, 64)
		if err != nil {
			tics.ThrowSys(Init, err)
		}

		if _, ok := idx.PathSet[el[0]]; ok {
			tics.ThrowSys(Init, duplFileErr(fpath, i))
		} else if !ok {
			idx.PathSet[el[0]] = i
		}

		e := Entry{
			NovemPath:   el[0],
			Inode:       inode,
			OutLinkPath: el[2],
			ChangedLast: el[3],
		}

		idx.Entries = append(idx.Entries, e)
	}

	return idx
}

func indexErr(idxFile string, lnNo int) error {
	s := fmt.Sprintf(
		"problem indexing file %v: line %v",
		idxFile, lnNo,
	)

	return errors.New(s)
}

func duplFileErr(idxFile string, lnNo int) error {
	s := fmt.Sprintf(
		"duplicate files in index '%v': line %v",
		idxFile, lnNo,
	)

	return errors.New(s)
}

func uninitWarn(fpath string) Index {
	wrn := tics.Make(" ?> novem index empty / not initialized at: ").Yellow()
	s := tics.Make(fpath).Bold()
	fmt.Printf("%v%v\n", wrn, s)
	return initEmptIdx()
}

func initEmptIdx() Index {
	return Index{
		Entries: []Entry{},
		PathSet: map[string]int{},
	}
}
