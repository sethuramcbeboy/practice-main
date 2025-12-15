package godevmostaskedcoding

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		sorted_v := Sort_func(v)
		m[sorted_v] = append(m[sorted_v], v)
	}
	arr := [][]string{}
	for _, v := range m {
		arr = append(arr, v)
	}
	return arr
}

func Sort_func(s string) string {
	a := strings.Split(s, "")
	sort.Strings(a)
	return strings.Join(a, "")
}

func main_anagram() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("Grouped Anagrams:", groupAnagrams(strs))
}
