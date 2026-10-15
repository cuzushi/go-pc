package main
//lvl3
import "fmt"

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	for i := 0; i < len(s); i++ {
		uppercase := s[i] >= 'A' && s[i] <= 'Z'
		nb := s[i] >= '0' && s[i] <= '9'
		if i == len(s)-1 && uppercase {
			return s
		}
		if uppercase && (s[i+1] >= 'A' && s[i+1] <= 'Z') && i < len(s)-1 {
			return s
		}
		if nb {
			return s
		}

	}
	var res []byte
	for i := 0; i < len(s); i++ {
		if i>0&&(s[i] >= 'A' && s[i] <= 'Z') && (s[i-1] >= 'a' && s[i-1] <= 'z') {
			res = append(res, '_',s[i])
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}

func main() {
	fmt.Println(CamelToSnakeCase("helloWorld"))

	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
