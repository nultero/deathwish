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

func NoFunctionErr() {
	fmt.Println(FINGER + "no function passed")
	quit()
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

func ThrowMultipleDiffOptsErr() {
	fmt.Print(redError() + "multiple diff options passed")
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

func UnrecognizedArgErr(desc string) {
	fmt.Print(redError() + "'" + desc + "' is not a valid argument, flag, or filepath")
	quit()
}

func ProblemGettingWorkingDir() {
	fmt.Println(redError() + "problem getting the current working dir")
	quit()
}
