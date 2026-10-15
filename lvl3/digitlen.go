package main
//lvl3
import "fmt"

func DigitLen(n, base int) int {
	if n < 0 {
		n = -n
	}
	if base > 36 || base < 2 {
		return -1
	}
	counter := 0
	for n != 0 {
		n = n / base
		counter++
	}
	return counter
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
