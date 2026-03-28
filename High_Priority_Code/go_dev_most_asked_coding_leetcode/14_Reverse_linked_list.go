package godevmostaskedcoding

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var s []int
	current := head
	for current != nil {
		s = append(s, current.Val)
		current = current.Next
	}
	var head2 *ListNode
	var current2 *ListNode
	for i := len(s) - 1; i >= 0; i-- {
		if head2 == nil {
			head2 = &ListNode{Val: s[i]}
			current2 = head2
		} else {
			current2.Next = &ListNode{Val: s[i]}
			current2 = current2.Next
		}
	}
	current2 = head2
	return current2
}

func main_reverse_linked_list() {
	var head *ListNode
	var current *ListNode
	for i := 1; i <= 10; i++ {
		if head == nil {
			head = &ListNode{Val: i}
			current = head
		} else {
			current.Next = &ListNode{Val: i}
			current = current.Next
		}
	}
	current = head
	for current != nil {
		//fmt.Println(current.Val)
		current = current.Next
	}
	current = head
	res := reverseList(current)
	for res != nil {
		fmt.Println("----", res.Val)
		res = res.Next
	}
}

//This one also reverse lined list in simple manner

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