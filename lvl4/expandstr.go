package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	var res string
	inword := false
	for _, c := range os.Args[1] {
		if c != ' ' {
			if !inword && res != "" {
				res += "   "
			}
			inword = true
			res += string(c)

		} else {
			inword = false
		}

	}
	for _, j := range res {
		z01.PrintRune(j)
	}
	z01.PrintRune('\n')
}
