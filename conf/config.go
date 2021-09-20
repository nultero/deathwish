package conf

import (
	"novem/nv"
	"os"
)

func Ok(path string) (string, string, error) {

	path = nv.GetFullPath(path)

	_, err := os.Stat(path)

	if err != nil {
		return "", "", err
	}

	config, r := os.ReadFile(path + "/novem.txt")

	if r != nil {
		return "", "", r
	}

	return string(config), path, nil
}
