package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	first := os.Args[1]
	sec := os.Args[2]

	m := map[rune]int{}
	for _, c := range first {
		m[c]++
		if m[c] == 1 {
			fmt.Print(string(c))
		}
	}
	for _, c := range sec {
		m[c]++
		if m[c] == 1 {
			fmt.Print(string(c))
		}
	}
	fmt.Println()
	return
}
