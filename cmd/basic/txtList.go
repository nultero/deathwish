package basic

import (
	"os"
	"strings"

	"github.com/nultero/tics"
)

func GetBasicList(dataDir string) []string {
	path := dataDir + "/basicNovLisr.txt"
	bytes, err := os.ReadFile(path)
	if err != nil {
		tics.ThrowSys(GetBasicList, err)
	}

	list := string(bytes)
	lines := strings.Split(list, "\n")
	return lines
}
