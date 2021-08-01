package main

import (
	"encoding/json"
	"fmt"
	"novem/bus"
	"os"
)

func Ls(lb *bus.Bus) {

	if len(*&lb.Paths) == 0 { // no paths, lists the json

		jsonF, err := os.ReadFile(novemJson())
		if err != nil {
			fmt.Println(err)
		}

		j, _ := json.Marshal(jsonF)

		fmt.Println(string(j))

	} else { // lists logged files in given PATHS
		for i := range *&lb.Paths {
			fmt.Println(*&lb.Paths[i])
		}
	}
}
