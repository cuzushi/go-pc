package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"
	for i, c := range arr {
		z01.PrintRune(rune(hex[c/16]))
		z01.PrintRune(rune(hex[c%16]))
		if (i+1)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		} else { 
			z01.PrintRune(' ')
		}

	} 
	for _, c := range arr {
		if c >=32 && c<=126 {
			z01.PrintRune(rune(c))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')

}

func mainwwww() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
