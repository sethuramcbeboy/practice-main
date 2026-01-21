package godevmostaskedcoding

import (
	"fmt"
)

func Move_all_zeros_to_end(s []int) []int{
	left := 0 
	for i:=0;i<len(s);i++{
		if s[i] != 0{
			s[left] = s[i]
			left++
		}
	}
	for left < len(s){
		s[left] = 0
		left++
	}
	return s
}

func main() {
	nums := []int{0,1,0,3,2,4}
	fmt.Println("Result is ", Move_all_zeros_to_end(nums))
}