package index

import "testing"

const (
	somePath     = "/somePath"
	someLongPath = "/home/maybeSomeUser/someotherpath/somegonkulous.txt"
)

var testPathSet = map[string]int{
	somePath:     4,
	someLongPath: 5,
}

var testIndex = Index{
	PathSet: testPathSet,
}

func TestHasFile(t *testing.T) {
	if !testIndex.HasFile("/somePath") {
		t.Errorf("wanted %v in test index, got: %v", somePath, testIndex.PathSet)
	}

	if !testIndex.HasFile("/home/maybeSomeUser/someotherpath/somegonkulous.txt") {
		t.Errorf("wanted %v in test index, got: %v", someLongPath, testIndex.PathSet)
	}
}
