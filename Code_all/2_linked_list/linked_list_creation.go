package practice

import "fmt"

func Create_Linked_list() *ListNode {
	var n, val int
	fmt.Println("Enter list size:")
	fmt.Scan(&n)
	if n == 0 {
		return nil
	}

	var head, current *ListNode
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		node := &ListNode{Val: val}
		if head == nil { // Corrected condition to initialize the head
			head = node
			current = head
		} else {
			current.Next = node
			current = current.Next
		}
	}
	return head
}