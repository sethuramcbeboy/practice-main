package practice

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // Create a dummy node to act as the start of the merged list
    dummy := &ListNode{}
    current := dummy

    // Traverse both lists and compare their values
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }

    // If there are remaining nodes in either list1 or list2, append them
    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }

    // Return the merged list starting from the first real node (skip the dummy node)
    return dummy.Next
}

func printList(head *ListNode) {
    for head != nil {
        fmt.Print(head.Val, " ")
        head = head.Next
    }
    fmt.Println()
}

func Call_list() {
    // Example usage
    list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
    list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

    mergedList := MergeTwoLists(list1, list2)
    printList(mergedList) // Output: 1 1 2 3 4 4
}
