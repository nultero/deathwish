package uvars

import (
	"deathwish/pathlib"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

// User variables.
type Uvars_t struct {
	Home          string
	DeathwishDir  string
	DeathwishFile string
	Dotfiles      Dots_t
}

func InitUvars() (Uvars_t, error) {
	u := Uvars_t{}

	home, err := os.UserHomeDir()
	if err != nil {
		return u, err
	}
	u.Home = home

	u.DeathwishDir = pathlib.SubCanon(deathwishDir, u.Home)
	u.DeathwishFile = pathlib.SubCanon(deathwishIndex, u.Home)

	dots := initDots()
	bytes, err := os.ReadFile(u.DeathwishFile)
	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			df := fmt.Sprintf("deathwish file \x1b[33m%s\x1b[0m does not exist", u.DeathwishFile)
			fmt.Println(ERRHEADER, df)

			fmt.Print("\n")

			fmt.Println(anywayWarn)

			err = os.Mkdir(u.DeathwishDir, 0755)
			if err != nil {
				fmt.Println(err)
			}

			_, err = os.Create(u.DeathwishFile)
			if err != nil {
				fmt.Println(err)
			}
			bytes = []byte{}

		} else {
			return u, err
		}
	}

	f := string(bytes)
	lines := strings.Split(f, "\n")

	for _, ln := range lines {
		if len(ln) != 0 {
			dots.Static[ln] = struct{}{}
		}
	}

	u.Dotfiles = dots

	return u, nil
}

// Saves the dotfiles to the index file.
func (u Uvars_t) Close() {

	if len(u.Dotfiles.NewDots) != 0 {
		u.meshAddedDots()
	}

}

func (u Uvars_t) meshAddedDots() {

	for _, s := range u.Dotfiles.NewDots {
		u.Dotfiles.Static[s] = struct{}{}
	}

	dots := []string{}

	for k := range u.Dotfiles.Static {
		if len(k) != 0 {
			dots = append(dots, k)
		}
	}

	sort.Strings(dots)
	dstr := strings.Join(dots, "\n")
	bytes := []byte(dstr)

	err := os.WriteFile(u.DeathwishFile, bytes, 0755)
	if err != nil {
		fmt.Printf(
			"%s problem writing to Deathwish index: %s\n%s",
			ERRHEADER,
			u.DeathwishFile,
			err,
		)
	}

	fmt.Print("\x1b[34m ADDED: \x1b[0m \n")
	for _, s := range u.Dotfiles.NewDots {
		fmt.Println(s)
	}
}
