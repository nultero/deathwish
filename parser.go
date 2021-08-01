package main

import (
	"fmt"
	"novem/bus"
	"novem/errs"
	"strings"
)

func parseArgs(args []string) {

	b := bus.NewBus()
	b.DataFile = novemJson()

	for !isEmpty(args) {

		r := args[0]

		if isFunc(r) {
			if isEmpty(b.Function) {
				b.Function = r
			} else {
				errs.ThrowMultipleFuncsErr()
			}

		} else if isFlag(r) {
			s := strings.ReplaceAll(r, "-", "")
			if len(s) == 1 {
				addFlag(s, &b)

			} else if len(s) == 0 {
				errs.EmptyFlagErr()

			} else { //multiple flags, "vvh" = verb +2, & help
				args = breakFlag(s, args)
			}

		} else if isDiffOpt(r) {
			if isEmpty(b.DiffOpts) {
				b.DiffOpts = r
			} else {
				errs.ThrowMultipleDiffOptsErr()
			}

		} //

		args = args[1:]
	}

	fmt.Println(b)
}

func addFlag(flag string, b *bus.Bus) {
	switch flag {
	case "v":
		*&b.Verbosity = max(*&b.Verbosity)

	case "h":
		*&b.Help = true

	case "d":
		*&b.Diff = true
	}
}

func breakFlag(flag string, args []string) []string {
	for i := range flag {
		s := "-" + string(flag[i]) //reappend dash for loop
		args = append(args, s)     //to recognize them as flags
	}
	// main loop will still chop off the 0th index multiflag in args slice
	return args
}

func max(v int) int {
	if v != 3 {
		return v + 1
	}

	return 3
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
