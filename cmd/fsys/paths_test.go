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

var meshSlash = []string{
	"/home/someUser/somedir/deeperDir/",
	"/file.ext",
	"/home/someUser/somedir/deeperDir/file.ext",
}

var meshAbsent = []string{
	"/home/someUser/somedir/deeperDir",
	"file.ext",
	"/home/someUser/somedir/deeperDir/file.ext",
}

func TestMeshPaths(t *testing.T) {
	dir, base := meshSlash[0], meshSlash[1]
	m := MeshPaths(dir, base)
	if m != meshSlash[2] {
		t.Errorf("wanted %v, got: %v", meshSlash[2], m)
	}

	dir, base = meshAbsent[0], meshAbsent[1]
	m = MeshPaths(dir, base)
	if m != meshAbsent[2] {
		t.Errorf("wanted %v, got: %v", meshAbsent[2], m)
	}

}
