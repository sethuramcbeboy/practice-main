package practice

import "fmt"


type Linkedlist struct{
	v int
	Next *Linkedlist
}

func linked_list_insert_at_middle() {
	var head, current *Linkedlist

	// creating liked list with value 0 -> 1 -> 2 -> 3 -> 4 -> nil
	for i := 0; i < 5; i++ {
		data := &Linkedlist{v: i}
		if head == nil {
			head = data
			current = head
		} else {
			for current.Next != nil {
				current = current.Next
			}
			current.Next = data
		}
	}



	// fiding lined list length
	current = head
	ctr := 0
	for current != nil {
		ctr++
		current = current.Next
	}
	// iterating the linked to middle 
	current = head
	for i := 0; i < ctr/2; i++ {
		current = current.Next
	}
	
	//new data
	//2200022 -> nil
	data := &Linkedlist{v: 220022}

	//binding the data at middle
	data.Next = current.Next
	//220022 -> 3 -> 4 -> nil
	current.Next = data
	// 0 -> 1-> 2 -> 220022 -> 3 -> 4 -> nil
	current = head

	for current != nil {
		fmt.Println(current.v)
		current = current.Next
	}
}
