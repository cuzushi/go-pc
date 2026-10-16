package main

import "fmt"

func FindPrevPrime(nb int) int {
	for n := nb; n >= 2; n-- {
		prime := true
		for i := 2; i*i <=n; i++ {
			if n%i == 0 {
				prime = false
				break
			}
		}
		if prime {
			return n
		}
	}
	return 0
}
func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
