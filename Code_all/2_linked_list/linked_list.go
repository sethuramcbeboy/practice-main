package practice

import "fmt"

// LinkedList represents a singly linked list
type LinkedList struct {
    head *Node
}


// Node represents a node in the linked list
type Node struct {
    value int
    next  *Node
}


// Insert adds a new node with the given value at the end of the list
func (ll *LinkedList) Insert(value int) {
    newNode := &Node{value: value}
    if ll.head == nil {
        ll.head = newNode
    } else {
        current := ll.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

// Delete removes the first occurrence of the node with the given value
func (ll *LinkedList) Delete(value int) {
    if ll.head == nil {
        return
    }
    // If the node to be deleted is the head
    if ll.head.value == value {
        ll.head = ll.head.next
        return
    }
    current := ll.head
    for current.next != nil && current.next.value != value {
        current = current.next
    }
    if current.next != nil {
        current.next = current.next.next
    }
}

// Display prints the values of the nodes in the list
func (ll *LinkedList) Display() {
    current := ll.head
    for current != nil {
        fmt.Print(current.value, " -> ")
        current = current.next
    }
    fmt.Println("nil")
}

func Linked_list() {
    ll := &LinkedList{}
    ll.Insert(10)
    ll.Insert(20)
    ll.Insert(30)

    fmt.Println("Linked List after insertions:")
    ll.Display() // Output: 10 -> 20 -> 30 -> nil

    ll.Delete(20)

    fmt.Println("Linked List after deletion of 20:")
    ll.Display() // Output: 10 -> 30 -> nil
}
