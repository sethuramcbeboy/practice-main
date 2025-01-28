package practice

func LinkedList_Last_elem_find(l *ListNode) int{
	current:=l
	for current.Next!=nil{
		current=current.Next
	}
	return current.Val
}