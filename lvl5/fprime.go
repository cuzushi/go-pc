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
	for i := 3; i*i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	nb, err := strconv.Atoi(os.Args[1])
	if err != nil || nb == 0 {
		return
	}
	first := true
	for i := 2; i <= nb; i++ {
		if nb%i == 0 {
			if isprime(i) {
				if first {
					fmt.Print(i)
					first = false
				} else {
					fmt.Print("*")
					fmt.Print(i)
				}
				nb = nb / i
				i = 2
			}
		}
	}
	fmt.Println()

}
