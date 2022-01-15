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
var IdxFile = "/.novem_index"

type index map[string]interface{}

func Get(dataDir string) (index, error) {
	path := dataDir + IdxFile
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return tics.ToJson(bytes), nil
}

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
