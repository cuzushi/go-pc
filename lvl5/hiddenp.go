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
	j := 0
	val := ""
	for i := 0; i < len(sec); i++ {
		if j < len(first) {
			if first[j] == sec[i] {
				val += string(first[j])
				j++

			}
		}
	}
	if val == first {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
	return

}
