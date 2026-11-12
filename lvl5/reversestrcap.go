package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, a := range args {
		for i := 0; i < len(a); i++ {
			c := a[i]
			if c >= 'A' && c <= 'Z' {
				c += 32
			}
			if c>='a'&& c<='z'{
				if i == len(a)-1 || a[i+1] == ' ' {
				c -= 32
				}
			}
			z01.PrintRune(rune(c))
		}
		z01.PrintRune('\n')

	}
}
