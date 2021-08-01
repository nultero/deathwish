package bus

type Bus struct {
	DataFile  string
	Function  string
	Verbosity int
	Paths     []string
	Help      bool
	Diff      bool
	DiffOpts  string
	Recursive bool
}

// Builds an empty logic bus for filtering arguments (like paths or flags) into.
func NewBus() Bus {
	return Bus{
		DataFile:  "",
		Function:  "",
		Verbosity: 0,
		Paths:     []string{},
		Help:      false,
		Diff:      false,
		DiffOpts:  "",
		Recursive: false,
	}
}
