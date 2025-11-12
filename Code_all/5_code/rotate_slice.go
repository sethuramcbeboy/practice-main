package practice

func Rotate(s []int,ctr int) []int{
    n:=len(s)
	ctr%=n
	s=append(s[ctr:],s[:ctr]...)
	return s
}