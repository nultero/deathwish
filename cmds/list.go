package cmds

import (
	"fmt"
	"novem/bus"
	"novem/chks"
	"strings"
)

func Ls(b *bus.Bus) {

	// if no paths queried, lists the dots being tracked
	// under the current working dir
	if chks.IsEmpty(*&b.Paths) {

		wd := strings.Split(cwd(), "/")

		for i := range wd {
			fmt.Println(wd[i])
		}

	} else { // lists logged files in given PATHS

		wd := cwd()

		for i := range *&b.Paths {
			p := *&b.Paths[i]
			p = strings.ReplaceAll(p, wd, "")
			fmt.Println(p)
		}

	}
}

// j, _ := os.ReadFile(*&b.DataFile)

// m := map[string]interface{}{}
// r := json.Unmarshal(j, &m)
// if r != nil {
// 	fmt.Println(r)
// }

// fmt.Println("len m ", len(m))
// for k, v := range m {

// 	dot, ok := v.(string)
// 	if ok {
// 		s := fmt.Sprintf("%s : %s", k, dot)
// 		fmt.Println(s)
// 		continue
// 	}

// 	testPanic, ok := v.(int)
// 	if !ok {
// 		fmt.Println("aww shit", testPanic)
// 	}
// 	// val, ok := v.(interface{})
// 	// if ok {
// 	// 	s := fmt.Sprintf("%s : %s", k, val)
// 	// 	fmt.Println(s)

// 	// } else {
// 	// 	fmt.Println("not okay")
// 	// }
