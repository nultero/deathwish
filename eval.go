package main

import (
	"fmt"
	"novem/bus"
)

func evalFuncs(b *bus.Bus) {

	switch *&b.Function {

	case "puts":
		Puts(*&b)

	case "list":
		Ls(*&b)

	default:
		fmt.Println("yeet")
	}
}
