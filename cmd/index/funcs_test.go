package index

import "testing"

const somePath = "/somePath"

var testPathSet = map[string]int{
	somePath: 4,
}

var testIndex = Index{
	PathSet: testPathSet,
}

func TestHasFile(t *testing.T) {
	if !testIndex.HasFile(somePath) {
		t.Errorf("wanted %v in test index, got: %v", somePath, testIndex.PathSet)
	}
}
