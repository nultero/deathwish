package main

import "strings"

type LogicBus struct {
	Function  string
	Verbosity int
	Paths     []string
	Help      bool
	Diff      bool
	ConfPath  string
}

func Increment(i string) string {
	i += "v"
	return i
}

func IsFlag(arg string) bool {
	if strings.Contains(arg, "-") {
		return true
	} else {
		return false
	}
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

	}

	if len(arg) > 1 {
		return "multi"
	} else if len(arg) == 0 {
		Throw("emptyFlag")
	}

	return arg
}

func BreakMultiFlag(flag string) []string {
	flag = strings.ReplaceAll(flag, "-", "")

	var flags []string

	for i := range flag {
		s := "-" + string(flag[i])
		flags = append(flags, s)
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

func IsNotEmpty(args []string) bool {
	if len(args) == 0 {
		return false
	}

	return true
}

func PopLastElement(flags []string) []string {
	return flags[:len(flags)-1]
}

func Reverse(flags []string) []string {

	// this was a macro in a plugin
	for i, j := 0, len(flags)-1; i < j; i, j = i+1, j-1 {
		flags[i], flags[j] = flags[j], flags[i]
	}

	return flags
}

func PassedFunction(fn string) bool {
	return !IsEmpty(fn)
}
