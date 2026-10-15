package checkpoint
//lvl2
func CheckNumber(arg string) bool {
 bool:= false 
 for _, c:= range arg{
	if c >='0' && c<='9'{
		bool=true 
	}
 }
 return bool
}
