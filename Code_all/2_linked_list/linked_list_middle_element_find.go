package practice

func LinkedList_Middle_elem_find(l *ListNode) int{
	current,ctr:=l,0
	for current!=nil{
		current=current.Next
		ctr+=1
	}
	s:=ctr/2
	current=l
	for i:=0;i<s;i++{
		if current!=nil{
			current=current.Next
		}
	}
	return current.Val
}