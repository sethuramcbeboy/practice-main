package godevmostaskedcoding

import (
	"fmt"
)

func firstUniqChar(s string) int {
	m := make(map[rune]int)

	for _, r := range s {
		m[r]++
	}

	for i, v := range s {
		if m[v] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("aabb"))
}