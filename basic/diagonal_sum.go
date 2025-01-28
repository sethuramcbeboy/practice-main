package practice

import "fmt"


func Diagonal_sum(req [][]int) int{
	l:=len(req)
	fmt.Println(l)
	sum:=0
	for i:=0;i<l;i++{
		sum+=req[i][i]
		fmt.Println("---",sum)
		if i!=(l-i-1){
			sum+=req[i][l-i-1]
			fmt.Println("===",sum)
		}
	}
	if l%2!=0{
       sum+=req[l-2][l-2]
	}
	return sum
}