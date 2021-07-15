package main

import (
	"fmt"
	"os"
)

func dumpPath(p chan string, newpath string) {
	p <- newpath
}

func Puts(lb *LogicBus) {

	paths := *&lb.Paths
	valpaths := make([]string, len(paths))

	for i := range paths {
		f, err := os.Stat(paths[i])
		os.IsNotExist(err)
		if err != nil {
			fmt.Printf("%s '%s' file does not exist, skipping this path \n", novemIcon, paths[i])
		} else {
			fmt.Println(f.ModTime())
			valpaths = append(valpaths, paths[i])
		}
	}

	p := make(chan string)
	tmp := []string{}
	for i := range valpaths {
		go dumpPath(p, valpaths[i])
		tmp = append(tmp, <-p)
	}

	fmt.Println(tmp)

}
