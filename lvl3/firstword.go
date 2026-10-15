package main

import (
	"fmt"
)

func FirstWord(s string) string {
	firstword:=""
	for i:=0;i<len(s);i++{
		if s[i]==' '{
			if firstword!=""{
				break
			}
		}else{
			firstword+=string(s[i])
		}
	}
	return firstword+"\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
