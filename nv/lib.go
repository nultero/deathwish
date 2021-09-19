package nv

import (
	"bufio"
	"os"
	"strings"
)

func GetInput() string {
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	return strings.ReplaceAll(s, "\n", "")
}
