package errs

import (
	"fmt"
	"os"
)

const (
	FINGER    = "\u261E  "
	novemIcon = "\u277e >"
)

func NoArgs() {
	fmt.Println(FINGER + "no args passed")
}

func redError() string {
	return "\033[31;1;4merror:\033[0m "
}

func quit() {
	fmt.Printf("\n%s exiting \n", novemIcon)
	os.Exit(1)
}

func ThrowMultipleFuncsErr() {
	fmt.Print(redError() + "multiple functional arguments passed")
	quit()
}

func EmptyFlagErr() {
	fmt.Print(redError() + "empty flag passed")
	quit()
}

func UnrecognizedFlagErr(desc string) {
	fmt.Print(redError() + "'" + desc + "' is an unrecognized flag")
	quit()
}
