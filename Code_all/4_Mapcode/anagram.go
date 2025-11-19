package mapcode

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

// Helper function to sort a string's characters
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// Function to group anagrams
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
		log.Println("=======",anagramMap)
	}

	// Convert map values to a 2D array
	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
		log.Println("------",group)
	}
	return result
}

func Anagram() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("Grouped Anagrams:", groupAnagrams(strs))
}
