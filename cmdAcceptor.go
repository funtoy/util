package util

import "fmt"

func CmdAcceptor(fm map[string]func()) {
	for {
		var str string
		fmt.Scanln(&str)
		if f, ok := fm[str]; ok {
			f()
		}
	}
}
