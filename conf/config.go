package conf

import (
	"novem/nv"
	"os"
)

func Ok(path string) (string, error) {

	path = nv.ExpandTilde(path)

	_, err := os.Stat(path)

	if err != nil {
		return "", err
	}

	config, r := os.ReadFile(path)

	if r != nil {
		return "", r
	}

	return string(config), nil
}
