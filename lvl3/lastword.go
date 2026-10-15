package main

import (
	"fmt"
)

func LastWord(s string) string {
	lastword:=""
	for i:=len(s)-1;i>=0;i--{
		if s[i]==' '{
			if lastword!=""{
				break
			}
		}else{
			lastword=string(s[i])+lastword
		}
	}
	return lastword+"\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
