package main

import (
	"fmt"
)

func EvalFuncs(lb *LogicBus) {

	switch *&lb.Function {

	case "puts":
		Puts(*&lb)

	case "list":
		Ls(*&lb)

	default:
		fmt.Println("yeet")
	}
}
