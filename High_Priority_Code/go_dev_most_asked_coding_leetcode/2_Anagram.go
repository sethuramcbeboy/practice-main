package godevmostaskedcoding

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram_using_map(s string, t string) bool {  
	if len(s) != len(t) {
		return false
	}
	s = strings.ToLower(s)
	t = strings.ToLower(t)

	m := make(map[rune]int)

	for _, v := range s {
		m[v]++  // s strings ->  each variables count
	}
	for _, v := range t {
		m[v]-- // t strings ->   each variables count subracting to check whether both contain same occurance
	}

	for _, v := range m {
		if v != 0 { 		// If any one contains any value means the occurance is mismatching
			return false
		}
	}
	return true
}

func main_anagram() {
	nums := "newyork"
	target := "yorkNew"
	fmt.Println("Result is ", isAnagram_using_map(nums, target))
}

func isAnagram(s string, t string) bool { // high time complexity 
	if len(s) != len(t) {
		return false
	}
	s = strings.ToLower(s)
	t = strings.ToLower(t)
	fmt.Println(s, t)
	sarr := strings.Split(s, "")
	tarr := strings.Split(t, "")
	sort.Strings(sarr)
	sort.Strings(tarr)
	return strings.Join(sarr,"") == strings.Join(tarr,"")
}
