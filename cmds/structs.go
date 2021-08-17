package cmds

type dotfile struct {
	Name, Timestamp string
}

type dir struct {
	Name string
	Dots []dotfile
}
