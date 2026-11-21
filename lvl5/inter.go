package main

import (
	"fmt"
	"os"
)

func main() {
	first := os.Args[1]
	sec := os.Args[2]
	if len(os.Args) != 3 {
		return
	}
	m := map[rune]int{}
	for _, c := range first {
		for _, v := range sec {
			if c == v {
				m[c]++
				if m[c] == 1 {
					fmt.Print(string(c))
				}
			}
		}
	}
	fmt.Print("\n")
	return
}
