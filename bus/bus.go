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

func getArgType(arg string) string {
	switch arg {
	case "puts":
		return "puts"
	}

	return "plop"
}

// func getFlagType(arg string) string {
// 	arg = strings.ReplaceAll(arg, "-", "")

// 	switch arg {
// 	case "v":
// 		return "verb"

// 	case "h", "help":
// 		return "help"

// 	case "d", "diff":
// 		return "diff"

// 	case "r":
// 		return "recursive"

// 	}

// 	if len(arg) > 1 {
// 		return "multi"
// 	} else if len(arg) == 0 {
// 		throwError("emptyFlag")
// 	}

// 	return arg
// }

// func popLastElement(args []string) []string {

// 	args = args[:len(args)-1]

// 	return args
// }

// func passedFunction(fn *string) bool {
// 	return !isEmpty(*fn)
// }
