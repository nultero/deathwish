package main

import (
	"novem/bus"
	"novem/chks"
	"novem/cmds"
	"novem/errs"
	"path/filepath"
	"strings"
)

func parseArgs(args []string) {

	b := bus.NewBus()
	b.DataFile = novemJson()

	for !chks.IsEmpty(args) {

		r := args[0]

		if chks.IsFunc(r) {
			if chks.IsEmpty(b.Function) {
				b.Function = r
			} else {
				errs.ThrowMultipleFuncsErr()
			}

		} else if chks.IsFlag(r) {
			s := strings.ReplaceAll(r, "-", "")
			if len(s) == 1 {
				addFlag(s, &b)

			} else if len(s) == 0 {
				errs.EmptyFlagErr()

			} else { //multiple flags, "vvh" = verb +2, & help
				args = breakFlag(s, args)
			}

		} else if chks.IsDiffOpt(r) {
			if chks.IsEmpty(b.DiffOpts) {
				b.DiffOpts = r
			} else {
				errs.ThrowMultipleDiffOptsErr()
			}

		} else if chks.IsValidPath(r) {
			p, _ := filepath.Abs(r)
			b.Paths = append(b.Paths, p)

		} else {
			errs.UnrecognizedArgErr(r)
		}

		args = args[1:]
	} // end args loop

	if chks.IsEmpty(b.Function) {
		errs.NoFunctionErr()
	} else {
		cmds.EvalFuncs(&b)
	}
}

func addFlag(flag string, b *bus.Bus) {
	switch flag {
	case "v":
		*&b.Verbosity = max(*&b.Verbosity)

	case "h":
		*&b.Help = true

	case "r":
		*&b.Recursive = true

	case "d":
		*&b.Diff = true
	}
}

func breakFlag(flag string, args []string) []string {
	for i := range flag {
		s := "-" + string(flag[i]) //reappend dash for the main loop
		args = append(args, s)     //to recognize them as flags
	}
	// main loop will still chop off the 0th index multiflag in args slice
	return args
}

func max(v int) int {
	if v != 3 {
		return v + 1
	}

	return 3
}
