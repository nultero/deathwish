package main

import (
	"fmt"
	"os"
	"strconv"
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

	return pth
}

func NovemJsonPath() string {
	return HandleTildePaths(
		PATH + "novem.json",
	)
}

func AnsiColorString(s string, r, g, b int) string {
	a := make([]string, 3)
	n := []int{r, g, b}
	for i := range n {
		a[i] = strconv.Itoa(n[i])
	}
	return fmt.Sprintf("\033[38;2;%s;%s;%sm%s\033[0m", a[0], a[1], a[2], s)
}
