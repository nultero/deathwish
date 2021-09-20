package conf

import (
	"errors"
	"fmt"
	"io/fs"
	"novem/colors"
	"novem/errs"
	"novem/nv"
	"os"
	"strings"
)

func Fix(path string) {

	_, err := os.Stat(path)

	path = nv.GetHomeDir() + nv.ExpandTilde(path)

	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			novemDirNotExists(path)
			initNovemFile(path)
		}
	}
}

func isValid(a string) bool {
	if strings.Contains(a, "y") || strings.Contains(a, "n") {
		return true
	}

	return false
}

const prompt = "create new novem dir? "

func novemDirNotExists(path string) {
	s := "'" + colors.Blue(path) + "'"
	fmt.Println("novem directory", s, "does not exist / was deleted \n ")

	pr := colors.Emph("[ y / n ]")

	fmt.Print(prompt, pr, " ")
	answer := strings.ToLower(nv.GetInput())

	for !isValid(answer) {
		fmt.Println("\n", colors.FingerRed(), colors.Red("expecting valid input"))
		fmt.Print(prompt, pr, " ")
		answer = strings.ToLower(nv.GetInput())
	}

	if strings.Contains(answer, "y") {
		createConf(path)

	} else {
		fmt.Println(colors.NovemNine(), "> sure thing, pardner")
		os.Exit(0)
	}

}

func createConf(dir string) {
	err := os.Mkdir(dir, 0755)

	if err != nil {
		errs.SysErr(err)
	}

	dir = "'" + colors.DarkBlue(dir) + "'"
	fmt.Println("created", dir)
}

func initNovemFile(path string) {

	ft, err := os.Stat(path)
	if err != nil {
		errs.SysErr(err)
	}

	path = path + "/novem.txt"

	now := nv.GetTimeStr(ft.ModTime())
	b := []byte(path + nv.SEPARATOR + now)

	r := os.WriteFile(path, b, 0755)
	if r != nil {
		errs.SysErr(r)
	}

	fmt.Println("wrote out", colors.DarkBlue(path), "as datafile")
}
