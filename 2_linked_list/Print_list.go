package practice

import "fmt"

func Print_list(l *ListNode){
	current:=l
	for current!=nil{
		fmt.Print(current.Val,"->")
		current=current.Next
	}
	fmt.Print("nil")
	fmt.Println("")
}