package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	novemIcon = "\u277e >"
	PATH      = "~/.nultero/novem/"
)

// var helpFlag = flag.Bool("h, help", false, "prints out help")
// var diffFlag = flag.Bool("d, diff", false, " prints out diff")
// var verb = flag.Int("v", 0, "increases the verbosity of all output")

// func init() {
// 	flag.BoolVar(diffFlag, "d", false, "prints out diff")
// 	flag.Func("v", "increases the verbosity of all output", func(v string) error {
// 		*verb++
// 		return nil
// 	})
// }

func main() {

	flag.Parse()

	// bus := LogicBus{ //defined in Bus.go
	// 	Function:  "",
	// 	Verbosity: verb,
	// 	Help:      helpFlag,
	// 	Diff:      diffFlag,
	// 	ConfPath:  PATH,
	// }

	ch := os.Args

	for i := range ch {
		fmt.Println(ch[i])
	}

	// flags := Reverse(flag.Args())

	// for IsNotEmpty(flags) {

	// 	thisFlag := flags[len(flags)-1]

	// 	if IsFlag(thisFlag) {
	// 		switch GetFlagType(thisFlag) {

	// 		// needs major cleanup

	// 		case "help":
	// 			bus.Help = true

	// 		case "diff":
	// 			bus.Diff = true

	// 		case "verb":
	// 			bus.Verbosity++

	// 		case "multi":
	// 			flags = BreakMultiFlag(thisFlag)
	// 			fmt.Println(flags)

	// 		default:
	// 			thisFlag = strings.ReplaceAll(thisFlag, "-", "")
	// 			ThrowDescriptive("unrecognizedFlag", thisFlag)
	// 		}

	// 	} else if IsFunc(thisFlag) {
	// 		if IsEmpty(bus.Function) {
	// 			bus.Function = thisFlag
	// 		} else {
	// 			Throw("multipleFunctionsErr")
	// 		}
	// 	}

	// 	flags = PopLastElement(flags)

	// } /// end main argparse loop

	// if PassedFunction(bus.Function) {
	// 	fmt.Println(" PLUNK ")
	// }

	// if *helpFlag {
	// 	fmt.Println(" halp")
	// }

	// if *diffFlag {
	// 	fmt.Println(" diff")
	// }
	// fmt.Println(novemIcon + " yes")
	// fmt.Println(bus)
}
