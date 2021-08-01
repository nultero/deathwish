package main

import (
	"strings"
)

type logicBus struct {
	DataFile  string
	Function  string
	Verbosity int
	Paths     []string
	Help      bool
	Diff      bool
	DiffOpts  string
	Recursive bool
}

// Builds an empty logic bus for filtering arguments (like paths or flags) into.
func defaultBus() logicBus {
	return logicBus{
		DataFile:  PATH,
		Function:  "",
		Verbosity: 0,
		Paths:     []string{},
		Help:      false,
		Diff:      false,
		DiffOpts:  "",
		Recursive: false,
	}
}

func getArgType(arg string) string {
	switch arg {
	case "puts":
		return "puts"
	}

	return "plop"
}

// func getFlagType(arg string) string {
// 	arg = strings.ReplaceAll(arg, "-", "")

// 	switch arg {
// 	case "v":
// 		return "verb"

// 	case "h", "help":
// 		return "help"

// 	case "d", "diff":
// 		return "diff"

// 	case "r":
// 		return "recursive"

// 	}

// 	if len(arg) > 1 {
// 		return "multi"
// 	} else if len(arg) == 0 {
// 		throwError("emptyFlag")
// 	}

// 	return arg
// }

func breakMultiFlag(flag string) []string {
	flag = strings.ReplaceAll(flag, "-", "")

	var flags []string

	for i := range flag {
		s := "-" + string(flag[i]) //reappend dash for loop
		flags = append(flags, s)   //to recognize them as flags
	}

	// plus a junk burner flag for the main() loop's pop
	return append(flags, "#####")
}

func popLastElement(args []string) []string {

	args = args[:len(args)-1]

	return args
}

func passedFunction(fn *string) bool {
	return !isEmpty(*fn)
}
