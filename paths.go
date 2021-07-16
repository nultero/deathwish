package main

import (
	"fmt"
	"os"
)

func HandleTildePaths(pth string) string {

	tilde := []byte("~")

	if pth[0] == tilde[0] {
		pth = pth[1:]
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		}

		pth = home + pth
	}

	fmt.Println(pth)
	return pth
}

func NovemJsonPath() string {
	return HandleTildePaths(
		PATH + "novem.json",
	)
}
