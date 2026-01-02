package godevmostaskedcoding

import "fmt"

func isValid(s string) bool {
	m := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	var stack []rune
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stack = append(stack, v)
		} else if closing, ok := m[v]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != closing {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main_paranthesis() {
	s := "{}[]()"
	fmt.Println(isValid(s))
}