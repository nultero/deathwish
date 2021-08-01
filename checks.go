package main

import (
	"strconv"
	"strings"
)

func isEmpty(s interface{}) bool {

	t, ok := s.(string)
	if ok {
		if len(t) == 0 {
			return true
		}
	}

	sl, ok := s.([]string)
	if ok {
		if len(sl) == 0 {
			return true
		}
	}

	return false
}

func isDiffOpt(arg string) bool {

	if len(arg) != 2 {
		return false
	}

	firstChar := string(arg[0])
	_, err := strconv.Atoi(firstChar)
	if err != nil {
		return false
	}

	secChar := string(arg[1])
	switch secChar {
	case "d", "m":
		return true

	default:
		return false
	}
}

func isFlag(arg string) bool {
	return strings.Contains(arg, "-")
}

func isFunc(arg string) bool {
	switch arg {
	case "puts", "list", "tree", "remove", "zipto", "unpack":
		return true
	}

	return false
}
