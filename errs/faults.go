package errs

import (
	"fmt"
	"novem/colors"
	"os"
)

func quit() {
	os.Exit(1)
}

func SysErr(r error) {
	s := colors.Red("error:")
	fmt.Println(s, r)
	panic(r)
	// quit()
}
