package practice

import "fmt"

// Helper function to create a cycle in the linked list for testing
// Helper function to create a cycle in the linked list for testing
func CreateCycle(head *ListNode, pos int) {
	if pos == -1 {
		// No cycle should be created
		return
	}

	current := head
	var cycleNode *ListNode
	count := 0
	for current.Next != nil {
		if count == pos {
			cycleNode = current
		}
		current = current.Next
		count++
	}

	// Link the last node to the cycle node to create a loop
	current.Next = cycleNode
}


func DetectCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	// Initialize two pointers: slow and fast
	slow, fast := head, head

	// Traverse the list
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Move slow pointer by one
		fast = fast.Next.Next // Move fast pointer by two

		// If they meet, there is a cycle
		if slow == fast {
			return true
		}
	}

	// If we exit the loop, there's no cycle
	return false
}



//after the createcycle we try to print the list using normal print_list func we wroted it infinites so we writing special func to print the list after createcycle for checking purpose.

func SafePrint(head *ListNode, limit int) {
	current := head
	count := 0

	for current != nil && count < limit {
		fmt.Print(current.Val, " -> ")
		current = current.Next
		count++
	}
	fmt.Println("...")
}