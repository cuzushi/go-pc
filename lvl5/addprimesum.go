package main

import (
	"fmt"
	"os"
	"strconv"
)

func isprime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 || os.Args[1] < "0" {
		fmt.Println(0)
		return
	}
	n, _ := strconv.Atoi(os.Args[1])
	res := 0
	for i := n; i >= 2; i-- {
		if isprime(i) {
			res += i
		}
	}
	fmt.Println(res)
	return
}
