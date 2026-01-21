package alation

// Given two strings S1 and S2, find all starting indices in S1 where an anagram of S2 begins.
// An anagram is a rearrangement of characters of a string using all original characters exactly once.

// Input : S1 = "cbdcbcacddbcan", S2 = "abcc"
// Output : [3, 4]

// Explanation
// Length of S2 = 4
// Substrings of length 4 in S1 are checked using a sliding window
// Matching anagram substrings:
// Index 3 → "cbca"
// Index 4 → "bcac"
// Both substrings are anagrams of "abcc".
import (
	"fmt"
	"sort"
	"strings"
)

// this is time complexity is O(n) and below main() we have the second approach which is simple but sort func take O(n2) and make total time complexity to o(n3) not prefered for interview
func findAnagrams(s1 string, s2 string) []int {
	result := []int{}
	if len(s1) < len(s2) {
		return result
	}
	need := make(map[byte]int)
	window := make(map[byte]int)
	// Frequency map for s2
	for i := 0; i < len(s2); i++ {
		need[s2[i]]++
	}
	left := 0
	windowSize := len(s2)
	for right := 0; right < len(s1); right++ {
		// Add current character to window
		window[s1[right]]++
		// Shrink window if size exceeds s2 length
		if right-left+1 > windowSize {
			window[s1[left]]--
			if window[s1[left]] == 0 {
				delete(window, s1[left])
			}
			left++
		}
		// Compare maps when window size matches
		if right-left+1 == windowSize && mapsEqual(window, need) {
			result = append(result, left)
		}
	}
	return result
}
// Helper function to compare two maps
func mapsEqual(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func main() {
	s1 := "cbdcbcacddbcan"
	s2 := "abcc"

	fmt.Println(findAnagrams(s1, s2))
}

// Very high time complexcity o(n3) due to inbuild sort func usage
func findAnagrams_simple(s1 string, s2 string) []int {
	var res []int
	if len(s1) < len(s2) {
		return res
	}
	for i := 0; i < len(s1)-len(s2); i++ {
		//using sliding sent 0 to 3 , 4 to 7 , 8 to 12
		if check(s1[i:len(s2)+i], s2) {
			res = append(res, i)
		}
	}
	return res
}

func check(s1, s2 string) bool {
	sarr := strings.Split(s1, "")
	s2arr := strings.Split(s2, "")
	sort.Strings(sarr)
	sort.Strings(s2arr)
	return strings.Join(sarr, "") == strings.Join(s2arr, "")
}
