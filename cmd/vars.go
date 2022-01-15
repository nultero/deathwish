package cmd

import (
	"novem/cmd/index"

	"github.com/nultero/tics"
)

var (
	NovNine    = tics.Make("\u277e").Bold().String()
	flavorText = tics.Make(" -> the ").Blue().String() +
		NovNine +
		tics.Make(" th time I've needed this").Blue().String()

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
	dataDir + ": " + index.IdxFile,
	idxFile + ": " + index.IdxFile,
}
