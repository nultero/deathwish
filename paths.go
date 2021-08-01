package main

import (
	"fmt"
	"novem/errs"
	"os"
	"strconv"
)

func cwd() string {
	d, err := os.Getwd()
	if err != nil {
		errs.ProblemGettingWorkingDir()
	}

	return d
}

func expandTildePath(p string) string {

	tilde := []byte("~")[0]

	if p[0] == tilde {
		p = p[1:]
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		}

		p = home + p
	}

	return p
}

func novemJson() string {
	return expandTildePath(PATH + "novem.json")
}

func ansiColorString(s string, r, g, b int) string {
	a := make([]string, 3)
	n := []int{r, g, b}
	for i := range n {
		a[i] = strconv.Itoa(n[i])
	}
	return fmt.Sprintf("\033[38;2;%s;%s;%sm%s\033[0m", a[0], a[1], a[2], s)
}
