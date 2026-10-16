package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args[1]
	a := os.Args[2]
	b := os.Args[3]
	if len(os.Args) != 4 {
		return
	}
	r := ""
	for _, v := range s {
		if string(v) == a {
			r += b

		} else {
			r += string(v)
		}
	}
	for _, d := range r {
		z01.PrintRune(d)

	}
	z01.PrintRune('\n')

}
