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

	CreateConfJsonIfNotExists()

	bus := LogicBus{ //defined in Bus.go
		ConfPath:  PATH,
		Function:  "",
		Verbosity: 0,
		Paths:     []string{},
		Help:      false,
		Diff:      false,
		DiffOpts:  "",
		Recursive: false,
	}

	args := Reverse(os.Args[1:])

	for IsNotEmpty(&args) {

		thisArg := args[len(args)-1]

		if IsFlag(thisArg) {

			switch GetFlagType(thisArg) {

			case "help":
				bus.Help = true

			case "diff":
				bus.Diff = true

			case "recursive":
				bus.Recursive = true

			case "verb":
				if bus.Verbosity < 3 {
					bus.Verbosity++
				} // ^ 3 is max verbosity anyway

			case "multi":
				args = PopLastElement(&args)
				r := BreakMultiFlag(thisArg)
				for i := range r {
					args = append(args, r[i])
				}

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
		} else {
			bus.Paths = append(bus.Paths, thisArg)
		}

		args = PopLastElement(&args)

	} /// end main argparse loop

	if bus.Help { // help is higher priority than other stuff
		if PassedFunction(&bus.Function) {
			EvalHelp(bus.Function)
		} else {
			PrintAllHelp()
		}

	} else if PassedFunction(&bus.Function) {
		EvalFuncs(&bus)

	} else {
		fmt.Println(novemIcon, "no args passed -> nothing doing")
		fmt.Println(novemIcon, "run `novem -h` for a list of arguments")
	}
}

func CreateConfJsonIfNotExists() {
	_, err := os.Stat(NovemJsonPath())
	if err != nil {
		fmt.Println(novemIcon, "the novem.json does not exist / been deleted")
		fmt.Printf(">>> create new dotfiles log? [ Y / n ] ")

		var in string
		_, r := fmt.Scanln(&in)

		if r != nil {
			fmt.Println(r)
			return
		}

		fmt.Println(novemIcon, in)
	}
}
