package index

const charSep = "»" // literally no one will use in a path, esp. not me, nope

type Entry struct {
	NovemPath   string
	Inode       uint64
	OutLinkPath string // the other link to this inode
	ChangedLast string
}
