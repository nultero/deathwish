package main

import (
	"fmt"
	"os"
)

func ThrowError(err string) {

	fmt.Print(novemIcon)

	switch err {
	case "multipleFunctionsErr":
		fmt.Print(" ERROR: multiple functional arguments passed")

	case "emptyFlag":
		fmt.Print(" ERROR: empty flag passed")

	}

	fmt.Printf("\n%s exiting \n", novemIcon)
	os.Exit(0)
}

func ThrowDescriptiveError(err string, desc string) {
	fmt.Print(novemIcon)

	switch err {
	case "unrecognizedFlag":
		fmt.Print(" ERROR: '" + desc + "' is an unrecognized flag")
	}

	fmt.Printf("\n%s exiting \n", novemIcon)
	os.Exit(0)
}
