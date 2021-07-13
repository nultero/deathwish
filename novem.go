package main

import (
	"fmt"
	"os"
	"strings"
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

	args := Reverse(os.Args[1:])

	for IsNotEmpty(args) {

		thisArg := args[len(args)-1]

		if IsFlag(thisArg) {
			switch GetFlagType(thisArg) {

			case "help":
				bus.Help = true

			case "diff":
				bus.Diff = true

			case "verb":
				if bus.Verbosity < 3 {
					bus.Verbosity++
				} // ^ 3 is max verbosity anyway

			case "multi":
				args = BreakMultiFlag(thisArg)

			default:
				thisArg = strings.ReplaceAll(thisArg, "-", "")
				ThrowDescriptiveError("unrecognizedFlag", thisArg)
			}

		} else if IsFunc(thisArg) {
			if IsEmpty(bus.Function) {
				bus.Function = thisArg
			} else {
				ThrowError("multipleFunctionsErr")
			}
		} else if IsAlphanumericArg(thisArg) {
			// placeholder
			fmt.Println("Haven't gotten to figuring how to validate this yet")
			// placeholder
		}

		args = PopLastElement(args)

	} /// end main argparse loop

	if PassedFunction(bus.Function) {
		fmt.Println(" PLUNK ")
	}

	fmt.Println(novemIcon + " yes")
	fmt.Println(bus)
}
