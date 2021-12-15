package cmd

import "github.com/nultero/tics"

var NovNine = tics.MakeT("\u277e").Bold().Str()
var flavorText = tics.Blue(" -> the ") + NovNine + tics.Blue(" th time I've needed this")

var SubDirFlag bool

var confFile = "confFile"
var dataDir = "dataDir"

var confMap = map[string]string{
	confFile: "$USER/.config/novem.yaml",
	dataDir:  "$USER/novem",
}

var defaultSettings = []string{ // TODO useful defaults around indexing, persists, sym v hardlinks?
}
