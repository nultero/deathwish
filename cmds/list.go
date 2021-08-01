package cmds

import (
	"encoding/json"
	"fmt"
	"novem/bus"
	"os"
)

func Ls(b *bus.Bus) {

	if len(*&b.Paths) == 0 { // no paths, lists the json

		jsonF, err := os.ReadFile(*&b.DataFile)
		if err != nil {
			fmt.Println(err)
		}

		j, _ := json.Marshal(jsonF)

		fmt.Println(string(j))

	} else { // lists logged files in given PATHS
		for i := range *&b.Paths {
			fmt.Println(*&b.Paths[i])
		}
	}
}
