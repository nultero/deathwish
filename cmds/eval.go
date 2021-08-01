package cmds

import (
	"novem/bus"
)

const novemIcon = "\u277e >"

func EvalFuncs(b *bus.Bus) {

	switch *&b.Function {

	case "puts":
		Puts(*&b)

	case "list":
		Ls(*&b)

	}
}
