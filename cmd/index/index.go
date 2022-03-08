package index

import (
	"encoding/json"
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

type index []Entry

// Reads index into memory, allows operations on
// idx. Hard crash if this fails; I want this to
// be obvious when this goes wrong.
func Init(fpath string) index {
	bytes, err := os.ReadFile(fpath)
	if err != nil {
		tics.ThrowSys(Init, err)
	}
	s := string(bytes)

	lines := strings.Split(s, "\n")
	idx := make([]Entry, len(lines))

	for i, ln := range lines {
		el := strings.Split(ln, charSep)
		if len(el) != 4 {
			tics.ThrowSys(Init, indexErr(fpath, i))
		}

		inode, err := strconv.Atoi(el[1])
		if err != nil {
			tics.ThrowSys(Init, err)
		}

		e := Entry{
			NovemPath:   el[0],
			Inode:       inode,
			OutLinkPath: el[2],
			ChangedLast: el[3],
		}

		idx[i] = e
	}

	return idx
}

// func Get(dataDir string) (index, error) {
// 	path := dataDir + IdxFile
// 	bytes, err := os.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return tics.ToJson(bytes), nil
// }

func From(fileNames []string, inodes []int) map[int]interface{} {
	var index = map[int]interface{}{}
	for i, node := range inodes {
		index[node] = fileNames[i]
	}

	return index
}

// func GetIndex(confMap map[string]interface{}) map[int]interface{} {

// 	dir, ok := confMap[]

// }

func WriteIndex(dataDir string, idx map[int]interface{}) {
	bytes, err := json.Marshal(idx)
	if err != nil {
		tics.ThrowSys(WriteIndex, err)
	}

	err = os.WriteFile(dataDir+IdxFile, bytes, tics.Perm)
	if err != nil {
		tics.ThrowSys(WriteIndex, err)
	}
}

func indexErr(idxFile string, lnNo int) error {
	s := fmt.Sprintf(
		"problem indexing file %v: line %v",
		idxFile, lnNo,
	)

	return errors.New(s)
}
