package main

import (
	"strconv"
	"strings"
)

type LogicBus struct {
	ConfPath  string
	Function  string
	Verbosity int
	Paths     []string
	Help      bool
	Diff      bool
	DiffOpts  string
	Recursive bool
}

func IsAlphanumericArg(arg string) bool {
	firstChar := string(arg[0])
	_, err := strconv.Atoi(firstChar)
	if err != nil {
		return false
	}

	return true
}

func IsFlag(arg string) bool {
	if strings.Contains(arg, "-") {
		return true
	} else {
		return false
	}
}

func GetArgType(arg string) string {
	switch arg {
	case "puts":
		return "puts"
	}

	return "plop"
}

func GetFlagType(arg string) string {
	arg = strings.ReplaceAll(arg, "-", "")

	switch arg {
	case "v":
		return "verb"

	case "h", "help":
		return "help"

	case "d", "diff":
		return "diff"

	case "r":
		return "recursive"

	}

	if len(arg) > 1 {
		return "multi"
	} else if len(arg) == 0 {
		ThrowError("emptyFlag")
	}

	return arg
}

func BreakMultiFlag(flag string) []string {
	flag = strings.ReplaceAll(flag, "-", "")

	var flags []string

	for i := range flag {
		s := "-" + string(flag[i]) //reappend dash for loop
		flags = append(flags, s)   //to recognize them as flags
	}

	// plus a junk burner flag for the main() loop's pop
	return append(flags, "#####")
}

func IsFunc(arg string) bool {
	switch arg {
	case "puts", "list", "tree", "remove", "zipto", "unpack":
		return true
	}

	return false
}

func IsEmpty(conf string) bool {
	if len(conf) == 0 {
		return true
	}

	return false
}

func IsNotEmpty(args *[]string) bool {
	if len(*args) == 0 {
		return false
	}

	return true
}

func PopLastElement(args *[]string) []string {

	*args = (*args)[:len(*args)-1]

	return *args
}

func Reverse(flags []string) []string {

	// this was a macro in a plugin
	for i, j := 0, len(flags)-1; i < j; i, j = i+1, j-1 {
		flags[i], flags[j] = flags[j], flags[i]
	}

	return flags
}

func PassedFunction(fn *string) bool {
	return !IsEmpty(*fn)
}
