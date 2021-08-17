package main

// chunko chickens
// chicken sammiches
// chumpkin pops

import (
	"novem/chks"
	"novem/errs"
	"os"
)

const (
	novemIcon = "\u277e >"
	PATH      = "~/.nultero/"
)

func main() {

	args := os.Args[1:]

	if !chks.IsEmpty(args) {
		parseArgs(args)

	} else {
		errs.NoArgs()
	}
}

// func createConfIfNotExists() {
// 	_, err := os.Stat(NovemJsonPath())
// 	if err != nil {
// 		fmt.Println(novemIcon, "the novem.json does not exist / been deleted")
// 		fmt.Printf(">>> create new dotfiles log? [ Y / n ]  >")

// 		in := ""
// 		_, r := fmt.Scanln(&in)

// 		if r != nil {
// 			if r.Error() != "unexpected newline" {
// 				fmt.Println(r)
// 				return
// 			}
// 		}

// 		if in != "n" {
// 			fmt.Println(novemIcon, "creating the novem.json in ", PATH)
// 			os.Create(NovemJsonPath())
// 		}
// 	}
// }
