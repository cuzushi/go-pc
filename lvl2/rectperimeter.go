package checkpoint
//lvl2
func RectPerimeter(w, h int) int {
	if h>0 && w>0{
		return 2*(h+w)
	}
	return -1
}
