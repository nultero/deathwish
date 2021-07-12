package main

import (
	"flag"
	"fmt"
)

const (
	novemIcon = "\u277e >"
	PATH      = "~/.nultero/novem/"
)

func main() {

	bus := LogicBus{ //defined in Bus.go
		Function:  "",
		Verbosity: 0,
		Help:      false,
		Diff:      false,
		ConfPath:  PATH,
	}

	flag.Parse()

	fl := flag.Args()
	for i := range fl {

		if IsFlag(fl[i]) {

			fmt.Println(fl[i], " is a flag")
		}

		if IsFunc(fl[i]) {

		}
	}

	fmt.Println(bus)

	fmt.Println(novemIcon + " yes")
}
