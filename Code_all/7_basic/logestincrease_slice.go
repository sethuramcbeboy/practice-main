package practice

func Longest_increase_in_arr_or_slice(req []int) int{
	ctl:=0
	n:=len(req)
    for i:=0;i<n-1;i++{
		if  req[i]<req[i+1]{
			ctl+=1
		}
	}
	if req[n-2]<req[n-1]{
		ctl+=1
	}
	return ctl
}