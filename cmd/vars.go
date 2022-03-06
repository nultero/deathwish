package cmd

import (
	"fmt"
	"novem/cmd/index"

	"github.com/nultero/tics"
)

var (
	novNine    = tics.Make("\u277e").Bold().String()
	flavorText = fmt.Sprintf("%v%v%v",
		tics.Make(" -> the ").Blue().String(),
		novNine,
		tics.Make(" th time I've needed this").Blue().String(),
	)

	RecurseFlag bool
	Verbosity   int = 0

	confFile = "confFile"
	dataDir  = "dataDir"
	idxFile  = "index file"
)

var confMap = map[string]string{
	confFile: "$USER/.config/novem.yaml",
	dataDir:  "$USER/novem",
	idxFile:  index.IdxFile,
}

var defaultSettings = []string{ // TODO useful defaults around indexing, persists, sym v hardlinks?
	dataDir + ": " + confMap[dataDir],
	idxFile + ": " + index.IdxFile,
}
