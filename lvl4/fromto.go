package main

import (
	"fmt"
)

func FromTo(from int, to int) string {
	if (from > 99 || from < 0) || (to > 99 || to < 0) {
		return "Invalid \n"
	}
	res := ""
	if from <= to {
		for i := from; i <= to; i++ {
			res += string(rune(i/10) + '0')
			res += string(rune(i%10) + '0')

			if i < to {
				res += ", "
			}
		}
	}
	if from > to {
		for i := from; i >= to; i-- {
			res += string(rune(i/10) + '0')
			res += string(rune(i%10) + '0')

			if i > to {
				res += ", "
			}
		}
	}
	return res + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
