package main

import "fmt"

func CanJump(slice []uint) bool {
	if len(slice) == 0 {
		return false
	}
	for i := 0; i < len(slice); {
		if i < len(slice) {
			if i == len(slice)-1 {
				return true

			} else if slice[i] == 0 {
				return false
			} else {
				i += int(slice[i])
				continue
			}
		}
	}
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
