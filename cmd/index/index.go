package index

import (
	"encoding/json"
	"os"

	"github.com/nultero/tics"
)

// The index file is, more or less, the declarative state of
// what the dots should look like. Accidental moves and deletions
// should crop up here, and it should be a little bit faster than
// trying to read a dir/dirs into memory.
var file = "/.novem_index"

func From(fileNames []string, inodes []int) map[int]interface{} {
	var index = map[int]interface{}{}
	for i, node := range inodes {
		index[node] = fileNames[i]
	}
	return index
}

// func GetIndex(dataDir string) map[int]interface{} {

// }

func WriteIndex(dataDir string, idx map[int]interface{}) {
	bytes, err := json.Marshal(idx)
	if err != nil {
		tics.ThrowSys(WriteIndex, err)
	}

	err = os.WriteFile(dataDir+file, bytes, tics.Perm)
	if err != nil {
		tics.ThrowSys(WriteIndex, err)
	}
}
