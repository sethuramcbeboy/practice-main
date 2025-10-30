package practice

func LinkedList_Reverse(l *ListNode) *ListNode {
	var prev *ListNode = nil
	current:=l

	for current!=nil{
		next:=current.Next
		current.Next=prev
		prev=current
		current=next
	}

	return prev
}