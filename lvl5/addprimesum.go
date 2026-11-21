package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] < "0" {
		fmt.Println()
		return
	}
	nb, _ := strconv.Atoi(os.Args[1])
	fmt.Print(nb)

	res := 0
	booleen := true
	for i := nb; i >= 2; i-- {
		for j := 2; j*j <= nb; j++ {
			if i%j == 0 {
				booleen = false
			}
		}
		if booleen {
			res += i
		}
	}
	fmt.Println(res)

}
