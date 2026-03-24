package ibm

import (
	"fmt"
	"strings"
)

// You are given a list of messages where each message contains multiple words separated by spaces. You are also given a target word.
// A message is considered "spam" if the target word appears at least 2 times in that message. Otherwise, it is considered "Not spam".
// Write a function to classify each message as "spam" or "Not spam".
// 📥 Input
// s1 = ["Hi Hi Hi", "Hello Hello"]
// s2 = "Hello"
// 📤 Output
// A mapping of each message to its classification:
// {
//   "Hi Hi Hi": "Not spam",
//   "Hello Hello": "spam"
// }

func Spam_check(s1 []string, s2 string) map[string]string {
	res := make(map[string]string)
	for _, v := range s1 {
		c := strings.Split(v, " ")
		if check(c, s2) {
			res[v] = "spam"
		} else {
			res[v] = "Not spam"
		}
	}
	return res
}

func check(s1 []string, s2 string) bool {
	ctr := 0
	for _, v := range s1 {
		if v == s2 {
			ctr++
		}
	}
	return ctr >= 2  //in array if words, contains more than or equal to 2 spams(s2)
}

func main() {
	s1 := []string{"Hi Hi Hi", "Hello Hello"}
	s2 := "Hello"
	fmt.Println(Spam_check(s1, s2))
}