package main

import "fmt"

func IsCapitalized(s string) bool {
	cap:=true 
	if len(s)==0{
		cap=false 
	}
	for i:=0;i<len(s);i++{
		if s[i]==' '&&i<len(s)-1{
			if s[i+1]>='a'&& s[i+1]<='z'{
				cap=false
			}
		}
	}
	return cap 
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
