package main
//lvl3/hashcode.go
import "fmt"

func HashCode(dec string) string {
	res := ""
	for _, char := range dec {
		hash := (int(char) + len(dec)) % 127
		if hash < 34 {
			hash += 33
		}
		res =res+string(rune(hash))
	}
	return res
}
func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
