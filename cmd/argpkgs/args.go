package argpkgs

import (
	"github.com/nultero/tics"
)

var PutsAliases = []string{"add", "link"}
var recursiveStr = tics.Make("-r").Blue().String()
var PutsDesc = "adds file arguments to the novem index" +
	" & hardlinks them in novem's dir. traverses" +
	" directories if recurse flag (" + recursiveStr + ") is set"
