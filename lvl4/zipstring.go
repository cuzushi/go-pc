package main

import (
	"fmt"
)

func ZipString(s string) string {
	res := ""
	r := 1
	for i := 0; i < len(s); i++ {
		if (s[i] > 'A' && s[i] < 'Z') || (s[i] > 'a' && s[i] < 'z') && i<len(s)-1 {
			if s[i] == s[i+1] {
				r++
				
			} else {
				res += string(rune('0'+r/1)) + string(s[i])
				r = 1
			}
		} else {
			res += string(s[i])
		}
	}
	return res
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
