package index

import "testing"

const (
	somePath     = "/somePath"
	someLongPath = "/home/maybeSomeUser/someotherpath/somegonkulous.txt"
)

var testPathMatches = map[string]int{
	somePath:     4,
	someLongPath: 5,
}

var testPathSet = map[string]struct{}{
	somePath:     {},
	someLongPath: {},
}

var testIndex = Index{
	PathMatches: testPathMatches,
	PathSet:     testPathSet,
}

func TestHasFile(t *testing.T) {
	if !testIndex.HasFile("/somePath") {
		t.Errorf("wanted %v in test index, got: %v", somePath, testIndex.PathSet)
	}

	if !testIndex.HasFile("/home/maybeSomeUser/someotherpath/somegonkulous.txt") {
		t.Errorf("wanted %v in test index, got: %v", someLongPath, testIndex.PathSet)
	}
}

// TODOO need to test match maps
