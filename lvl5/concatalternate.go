package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	if len(slice1)==len(slice2){
		return append(slice1,slice2...)
	}
	first := slice1
	sec := slice2
	if len(slice2) > len(slice1) {
		first = slice2
		sec = slice1
	}
	var res[]int 
	for i:=0;i<len(first);i++{
		if i<len(first){
			res=append(res, first[i])
		}
		if i<len(sec){
			res=append(res, sec[i])
		}
	}
	return res 

}
func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
