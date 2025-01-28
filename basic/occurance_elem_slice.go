package practice

func Occourance(req []int)map[int]int{
	ctr:=make(map[int]int)
	for _,v:=range req{
		ctr[v]++
	}
	return ctr
}