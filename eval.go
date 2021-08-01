package main

import (
	"fmt"
)

func evalFuncs(b *logicBus) {

	switch *&b.Function {

	case "puts":
		Puts(*&b)

	case "list":
		Ls(*&b)

	default:
		fmt.Println("yeet")
	}
}
