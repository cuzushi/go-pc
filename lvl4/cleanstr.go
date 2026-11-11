package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	res := ""
	inword := true
	for _, c := range os.Args[1] {
		if c != ' ' && c != '\t' {
			if !inword && len(res)>0{
				res+=" "
			}
			res += string(c)
			inword = true
		} else {
			inword=false
		}
	}
	for _, c := range res {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
