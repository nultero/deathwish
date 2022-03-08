package fsys

import "testing"

const homeTest = "/home/someUser"

var pathTestCases = []string{
	"/home/someUser/somedir/file.ext",
	"/home/someUser/somedir/deeperDir/otherfile",
}

var pathTrailCases = []string{
	"somedir",
	"somedir/deeperDir",
}

var pathBaseCases = []string{
	"file.ext",
	"otherfile",
}

func TestTrail(t *testing.T) {
	for i, c := range pathTestCases {
		trail, base := Trail(c, homeTest)
		if trail != pathTrailCases[i] {
			t.Errorf("wanted %v, got: %v", pathTrailCases[i], trail)
		}

		if base != pathBaseCases[i] {
			t.Errorf("wanted %v, got: %v", pathBaseCases[i], base)
		}
	}
}
