package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	first := []rune(os.Args[1])
	sec := []rune(os.Args[2])
	i := 0
	for _, c := range sec {
		if i < len(first) && c == first[i] {
			i++
		}
	}
	if i == len(first) {
		for _, c := range first {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}

}
