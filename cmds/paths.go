package cmds

import (
	"novem/errs"
	"os"
)

func cwd() string {
	d, err := os.Getwd()
	if err != nil {
		errs.ProblemGettingWorkingDir()
	}

	return d
}
