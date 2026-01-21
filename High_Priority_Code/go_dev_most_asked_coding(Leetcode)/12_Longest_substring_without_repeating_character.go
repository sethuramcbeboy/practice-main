package godevmostaskedcoding

import "fmt"

//Dynamic Sliding window 
func lengthOfLongestSubstring(s string) int {
	left := 0
	lastIndex := make(map[rune]int)
    maxLen := 0

    for right, ch := range s {
        if idx, ok := lastIndex[ch]; ok && idx >= left {
            left = idx + 1
        }
        lastIndex[ch] = right
        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }
    return maxLen
}

func main() {
	s := "abcabcab"
	fmt.Println(lengthOfLongestSubstring(s))
}