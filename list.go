package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Ls(lb *LogicBus) {

	if len(*&lb.Paths) == 0 { // no paths, lists the json

		noveHome := NovemJsonPath()

		jsonF, err := os.ReadFile(noveHome)
		if err != nil {
			fmt.Println(err)
		}

		j, _ := json.Marshal(jsonF)

		fmt.Println(j)

	} else { // lists logged files in given PATHS
		for i := range *&lb.Paths {
			fmt.Println(*&lb.Paths[i])
		}
	}
}
