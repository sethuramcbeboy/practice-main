package leetcode 

import "log"

type ListNode struct {
     Val int
     Next *ListNode
}

func Leet_83(){
	arr:=[]int{1,1,2,3,4,5,5}
	head:=&ListNode{Val: arr[0]}
	current:=head
	for _,val:=range arr[1:]{
		current.Next=&ListNode{Val: val}
		current=current.Next
	}
	res:=deleteDuplicates(head)
	current=res
	for current!=nil{
		log.Println(current.Val,"<-")
		current=current.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head==nil{
		return nil
	}
	current:=head
	for current.Next!=nil{
		if current.Val==current.Next.Val{
			current.Next=current.Next.Next
			
		}else{
			current=current.Next
		}
	}
	return head
}

/*83. Remove Duplicates from Sorted List
Solved
Easy
Topics
Companies
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

 

Example 1:
Input: head = [1,1,2]
Output: [1,2]
Example 2:
Input: head = [1,1,2,3,3]
Output: [1,2,3]
*/