package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {
	res := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			for i := 0; i < int(c-'A')+1; i++ {
				res += string(c)

			}
		}
		if c >= 'a' && c <= 'z' {
			for i := 0; i < int(c-'a')+1; i++ {
				res += string(c)

			}
		}
		res += string(c)
	}
	return res
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
