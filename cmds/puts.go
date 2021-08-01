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
	if len(paths) == 0 { //early return if no paths
		fmt.Println(novemIcon, "no path args passed in to put into novem logs")
		return
	}

	validPaths := []string{}

	for i := range paths {
		_, err := os.Stat(paths[i])
		os.IsNotExist(err)
		if err != nil {
			s := "\033[31;1;4m" + paths[i] + "\033[0m"
			fmt.Printf("%s '%s' file does not exist, skipping this path \n", novemIcon, s)
		} else {
			// fmt.Println(f.ModTime())
			validPaths = append(validPaths, paths[i])
		}
	}

	tmp := []string{}

	for !chks.IsEmpty(validPaths) {

		thisPath := validPaths[len(validPaths)-1]

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

		// validPaths = popLastElement(validPaths)
	}

	fmt.Println(tmp)

}
