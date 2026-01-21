package godevmostaskedcoding

import (
	"fmt"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left := 0
	count := len(t)
	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		if need[s[right]] > 0 {
			count--
		}
		need[s[right]]--

		for count == 0 {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			need[s[left]]++
			if need[s[left]] > 0 {
				count++
			}
			left++
		}
	}

	if minLen > len(s) {
		return ""
	}
	return s[start : start+minLen]
}

func main_min_window() {
	s, t := "ADOBECODEBANC", "ABC"
	fmt.Println(minWindow(s, t))
}