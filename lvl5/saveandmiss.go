package main

import "fmt"

func SaveAndMiss(arg string, num int) string {
	val := ""
	if num <= 0 {
		return arg
	}
	for i := 0; i < len(arg); i++ {
		val += string(arg[i])

		if (i+1)%num == 0 {
			i += num

		}

	}
	return val
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}
