package main

import (
	"fmt"
	"novem/errs"
)

func parseArgs(args []string) {
	b := defaultBus()

	for i := range args {
		r := args[i]

		if isFunc(r) {
			if isEmpty(b.Function) {
				b.Function = r
			} else {
				errs.ThrowMultipleFuncsErr()
			}

		}
	}

	fmt.Println(b)
}

// for isNotEmpty(args) {

// 	thisArg := args[len(args)-1]

// 	if isFlag(thisArg) {

// 		switch GetFlagType(thisArg) {

// 		case "help":
// 			bus.Help = true

// 		case "diff":
// 			bus.Diff = true

// 		case "recursive":
// 			bus.Recursive = true

// 		case "verb":
// 			if bus.Verbosity < 3 {
// 				bus.Verbosity++
// 			} // ^ 3 is max verbosity anyway

// 		case "multi":
// 			args = PopLastElement(args)
// 			r := BreakMultiFlag(thisArg)
// 			for i := range r {
// 				args = append(args, r[i])
// 			}

// 		default:
// 			thisArg = strings.ReplaceAll(thisArg, "-", "")
// 			ThrowDescriptiveError("unrecognizedFlag", thisArg)
// 		}

// 	} else if IsFunc(thisArg) {
// 		if IsEmpty(bus.Function) {
// 			bus.Function = thisArg

// 		} else {
// 			ThrowError("multipleFunctionsErr")
// 		}

// 	} else if IsAlphanumericArg(thisArg) {
// 		// placeholder
// 		fmt.Println("Haven't gotten to figuring how to validate this yet")
// 		// placeholder
// 	} else {
// 		bus.Paths = append(bus.Paths, thisArg)
// 	}

// 	args = PopLastElement(args)

// } /// end main argparse loop

// if bus.Help { // help is higher priority than other stuff
// 	if PassedFunction(&bus.Function) {
// 		EvalHelp(bus.Function)
// 	} else {
// 		PrintAllHelp()
// 	}

// } else if PassedFunction(&bus.Function) {
// 	EvalFuncs(&bus)

// } else {
// 	fmt.Println(novemIcon, "no args passed -> nothing doing")
// 	fmt.Println(novemIcon, "run `novem -h` for a list of arguments")
// }
