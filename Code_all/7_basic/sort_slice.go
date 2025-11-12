package practice

import "fmt"

func Sort(){
	var s[]int
	s=append(s, 1,2,3,4,5)
	for i:=0;i<len(s)-1;i++{
		if s[i]>s[i+1]{
			fmt.Println("Unsorted array")
		    return
		}
	}
	fmt.Println("Sorted array")
}