package nv

import (
	"fmt"
	"novem/errs"
	"os"
	"strings"
	"time"
)

const rfc2822 = "Mon Jan 02 15:04:05 -0700 2006"
const HOME = "$HOME"

type NovemNode struct {
	PathChunks []string
	Timestamp  string
	SubNodes   *[]NovemNode
	Next       *NovemNode
}

type NovemFile struct {
	Filepath  string
	Timestamp string
}

type NovemDir struct {
	Filepath string
	Files    []NovemFile
	Dirs     []NovemDir
}

func (nd NovemDir) addFile(filepath string) {
	mt, err := os.Stat(filepath)
	if err != nil {
		errs.SysErr(err)
	}

	t, r := time.Parse(rfc2822, mt.ModTime().String())
	if r != nil {
		errs.SysErr(r)
	}

	nvf := NovemFile{
		filepath,
		t.String(),
	}

	nd.Files = append(nd.Files, nvf)
}

func (nd NovemDir) AddDir(filepath string) {

	files, err := os.ReadDir(filepath)
	if err != nil {
		errs.SysErr(err)
	}

	for file := range files {
		fmt.Println("in ", filepath, ":::", file)
	}
}

func (root NovemDir) ToString() string {

	return HOME
}

func (nn NovemNode) From(path string) NovemNode {

	rePath := strings.ReplaceAll(path, GetHomeDir(), "")
	spl := strings.Split(rePath, "/")
	nn.PathChunks = []string{}

	for _, s := range spl {
		if len(s) > 0 {
			fmt.Println("s: ", s)
			nn.PathChunks = append(nn.PathChunks, s)
		}
	}

	return nn
}

// Takes /home/$user/... and swaps $user with the static string "$HOME", for portability among paths.
func swapHomes(path string) string {
	return strings.ReplaceAll(path, GetHomeDir(), HOME)
}
