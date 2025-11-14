package practice

import "fmt"

func Max(){
	var s[]int
	max:=0
	s=append(s, 1,2,3,4,5,6,7,8)
	for _,v:=range s{
        if v>max{
			max=v
		}
	}
	fmt.Println(max)
}