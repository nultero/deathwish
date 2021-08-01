package cmds

type dotfile struct {
	name, timestamp string
}

type dir struct {
	name string
	dots []dotfile
}
