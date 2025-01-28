package practice

func LinkedList_insert_at_begining(v int, l *ListNode) *ListNode {

	current := &ListNode{Val: v}
	current.Next=l
	return current
}
