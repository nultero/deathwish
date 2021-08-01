package chks

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func IsEmpty(s interface{}) bool {

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

func IsDiffOpt(arg string) bool {

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

func IsFlag(arg string) bool {
	return strings.Contains(arg, "-")
}

func IsFunc(arg string) bool {
	switch arg {
	case "puts", "list", "tree", "remove", "zipto", "unpack":
		return true
	}

	return false
}

func IsValidPath(arg string) bool {
	p, r := filepath.Abs(arg)
	if r != nil {
		return false
	}

	_, err := os.Stat(p)
	if err != nil {
		return false
	}

	return true
}
