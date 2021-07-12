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

func Count(arg string, char string) int {
	return 0
}

func IsFlag(arg string) bool {
	if strings.Contains(arg, "-") {
		return true
	} else {
		return false
	}
}

func IsFunc(arg string) bool {
	switch arg {
	case "puts", "list", "tree", "remove", "zipto", "unpack":
		return true
	}

	return false
}
