package practice

import "fmt"

func LinkedList_insert_at_end(v int, list *ListNode) *ListNode {
	fmt.Println("insert func()",v)
	printList(list)
	newnode := &ListNode{Val: v}
	if list == nil {
		return newnode
	}
	current := list
	for current.Next != nil {
		current = current.Next
	}
	current.Next=newnode

	return list
}
