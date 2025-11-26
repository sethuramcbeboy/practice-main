package godevmostaskedcoding

import (
	"fmt"
	"log"
)

func check(s1, s2 string) int {
	log.Println("s1 s2",s1,s2)
	a,b,l := []rune(s1),[]rune(s2),len(s1)
	if len(s1) > len(s2) {
		l = len(s2)
	}
	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return l
}

func LongestCommonPrefix(s []string) int {
	prefix := s[0]
	min := len(s[0])
	for i := 1; i < len(s); i++ {
		res := check(prefix, s[i])
		if res < min {
			min = res
		}
		log.Println(res)
	}
	return min
}

func main() {
	s := []string{"flower", "flow", "flight"}
	prefix := s[0]
	res := LongestCommonPrefix(s)
	fmt.Println(prefix[:res])
}