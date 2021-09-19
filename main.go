package main

import (
	"novem/cmds"
	"novem/conf"
	"os"
)

const PATH = "~/.novem/"

func main() {

	//
	//
	//
	//
	//

	args := os.Args[1:]

	config, err := conf.Ok(PATH)

	if err != nil {
		conf.Fix(PATH)
	}

	cmds.Parse(args, config)
}
