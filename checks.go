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

func isAlphanumericArg(arg string) bool {
	firstChar := string(arg[0])
	_, err := strconv.Atoi(firstChar)
	if err != nil {
		return false
	}

	return true
}

func isFlag(arg string) bool {
	if strings.Contains(arg, "-") {
		return true
	} else {
		return false
	}
}

func isFunc(arg string) bool {
	switch arg {
	case "puts", "list", "tree", "remove", "zipto", "unpack":
		return true
	}

	return false
}
