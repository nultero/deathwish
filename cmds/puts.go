package cmds

import (
	"fmt"
	"novem/bus"
	"novem/chks"
	"os"
	"path/filepath"
)

func Puts(b *bus.Bus) {

	paths := *&b.Paths

	if chks.IsEmpty(paths) { //early return if no paths
		fmt.Println(novemIcon, "no path args passed in to put into novem log")
		return
	}

	tmp := []string{}

	for !chks.IsEmpty(paths) {

		thisPath := paths[0]

		f, _ := filepath.Abs(filepath.Join(thisPath))

		inf, _ := os.Stat(f)
		if inf.IsDir() {
			if *&b.Recursive {
				fmt.Println("recurse works")

			} else { // gotta clean up these ansis
				s := "\033[36;1;4m" + thisPath + "\033[0m"
				fmt.Println("... skipping", s, ": is a dir (no recursive flag found)")
			}

		} else {
			// plus := "\033[32m+\033[0m"
			// s := "\033[;0;4m" + AnsiColorString(thisPath, 50, 150, 50) + "\033[0m"
			// fmt.Println(plus, s)
			fmt.Println(inf)
			tmp = append(tmp, f)
		}

		paths = paths[1:]
	}

	fmt.Println(tmp)

}
