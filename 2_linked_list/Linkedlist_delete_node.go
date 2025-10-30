package practice

func LinkedList_delete_node(v int, l *ListNode) *ListNode {
	if l==nil{
		return nil
	}
	if l.Val==v{
		return l.Next
	}
	current := l
	for current.Next.Val != v && current.Next != nil {
		current = current.Next
	}
	if current.Next!=nil && current.Next.Val==v{
		current.Next=current.Next.Next
	}
	return l
}
