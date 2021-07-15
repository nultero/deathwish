package main

import (
	"fmt"
	"strings"
)

func BlankLine() {
	s := "\n\n"
	fmt.Printf("%s", s)
}

func EvalHelp(h string) {
	s := novemIcon + "    HELP"
	fmt.Println(s)
	fmt.Println(strings.Repeat("  -", len(s)/2))
	switch h {
	case "puts":
		PutsHelp()

	case "list":
		LsHelp()

	case "remove":
		RemoveHelp()

	case "unpack":
		UnpackHelp()

	case "zipto":
		ZiptoHelp()
	}
}

func PutsHelp() {
	fmt.Println(novemIcon, "  novem PUTS [ PATH ]")
	fmt.Println("\t        : logs the given PATH(s) as a dotfile")
	fmt.Println("\t        : (this will also overwrite an existing dot if novem already logged it)")
	BlankLine()
}

func LsHelp() {
	fmt.Println(novemIcon, "  novem LIST [ path ]")
	fmt.Println("\t         : lists all dotfiles, details depend on verbosity")
	fmt.Println("\t         : uses current dir if no specific path given")
	fmt.Println("\t{string} : searches for {str} in dotfiles")
	BlankLine()
}

func RemoveHelp() {
	fmt.Println(novemIcon, "  novem REMOVE [ PATH ]")
	fmt.Println("\t        : removes PATH from dotfiles")
	fmt.Println("\t          (after confirmation)")
	BlankLine()
}

func UnpackHelp() {
	fmt.Println(novemIcon, "  novem UNPACK [ .zip file ]")
	fmt.Println("\t          : attempts to unpack a .zip into a Novem json")
	fmt.Println("\t     (if valid, unpacks all dotfiles to their respective PATHs)")
	BlankLine()
}

func ZiptoHelp() {
	fmt.Println(novemIcon, "  novem ZIPTO  (optional zip output path)")
	fmt.Printf("\t>>> DEFAULT NOVEM ZIP PATH:  %s", PATH)
	BlankLine()
}

func PrintAllHelp() {
	PutsHelp()
	LsHelp()
	RemoveHelp()
	UnpackHelp()
	ZiptoHelp()
}
