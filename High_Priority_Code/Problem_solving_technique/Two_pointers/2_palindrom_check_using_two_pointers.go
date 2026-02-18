package twopointers

import "fmt"
// palindrome means reverse order is same or not
func Palindrome_check_using_two_pointers(s string) bool {
	s_rune := []rune(s)
	j := len(s_rune) - 1
	for i := 0; i < len(s_rune)/2; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

func main_s() {
	s := "abcdbbdcba"
	fmt.Println("String is palindrome", Palindrome_check_using_two_pointers(s))
}