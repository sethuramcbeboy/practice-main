package practice

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	stack1 Stack
	stack2 Stack
}

type Stack struct {
	elements []int
}

func (q *Queue) Enqueue(x int) {
	q.stack1.Push(x)
}

func (s *Stack) Push(x int) {
	s.elements = append(s.elements, x)
}

func (q *Queue) Dequeue() int {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			q.stack2.Push(q.stack1.Pop())
		}
	}
	return q.stack2.Pop()
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1
	}
	lastIndex := len(s.elements) - 1
	elem := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return elem
}

func Queue_2_stack_impl() {
	queue := Queue{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter operations (enqueue x or dequeue), followed by 'exit' to stop:")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		inputFields := strings.Fields(input)

		if len(inputFields) == 0 {
			fmt.Println("Empty input")
			continue
		}

		command := inputFields[0]

		switch command {
		case "enqueue":
			if len(inputFields) == 2 {
				value, err := strconv.Atoi(inputFields[1])
				if err != nil {
					fmt.Println("Error in given data:", err)
					continue
				}
				queue.Enqueue(value)
			} else {
				fmt.Println("Invalid enqueue command format. Usage: enqueue x")
			}
		case "dequeue":
			value := queue.Dequeue()
			if value != -1 {
				fmt.Println(value)
			} else {
				fmt.Println("Queue is empty")
			}
		case "exit":
			return
		default:
			fmt.Println("Invalid command:", input)
		}
	}
}
