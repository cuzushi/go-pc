package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	res := ""
	if n < 0 {
		s = "-"
		n = -n
	}
	for n != 0 {
		res = string(rune(n%10)+'0') + res
		n = n / 10
	}
	return s + res
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
