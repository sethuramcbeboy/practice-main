package stackqueue

import (
	"fmt"
)

type stack struct {
	v []int
}

func push(a int, s *stack) {
	s.v = append(s.v, a)
	fmt.Println(s.v)
}

func pop(s *stack) {
	if len(s.v) <= 0 {
		return
	}
	s.v = s.v[:len(s.v)-1]
	fmt.Println(s.v)
}

func main() {
	s := &stack{}
	push(1, s)
	push(2, s)
	push(3, s)
	push(4, s)
	pop(s)
	pop(s)
	pop(s)
	pop(s)
}
