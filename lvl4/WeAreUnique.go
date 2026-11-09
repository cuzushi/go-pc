package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	rune := make(map[rune]int)
	for _, r := range str1 {
		rune[r] |= 1
	}
	for _, r := range str2 {
		rune[r] |= 2

	}
	count := 0
	for _, r := range rune {
		if r == 1 || r == 2 {
			count++
		}
	}
	return count
}

func main11() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
