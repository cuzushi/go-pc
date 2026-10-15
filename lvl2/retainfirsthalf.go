package main
//lvl2
import "fmt"

func RetainFirstHalf(str string) string {
	return str[:(len(str)/2)+(len(str)%2)]
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
