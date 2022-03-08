package index

const charSep = "»"

type Entry struct {
	NovemPath   string // the novem variant of the path given to this entry
	Inode       int
	OutLinkPath string // the other link to this inode
	ChangedLast string
}
