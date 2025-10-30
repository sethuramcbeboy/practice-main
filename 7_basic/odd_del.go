package practice

import "fmt"

func Odd(){
	var s[]int
	j:=0
	s=append(s, 1,2,3,4,5,6,7,8,9,10)
	for i,_:=range s{
		if s[i]%2==0{
			s[j]=s[i]
			j++
		}
	}
    s=s[:j]
	fmt.Println(s,j)
}