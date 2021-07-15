package main

import (
	"fmt"
)

func dumpPath(p chan string, newpath string) {
	p <- newpath
}

func EvalFuncs(lb *LogicBus) {

	p := make(chan string)
	tmp := []string{}

	switch *&lb.Function {
	case "puts":
		paths := *&lb.Paths
		for i := range paths {
			go dumpPath(p, paths[i])
			tmp = append(tmp, <-p)
		}

	default:
		fmt.Println("yeet")
	}

	fmt.Println(tmp)
}
