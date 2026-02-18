package twopointers

import "fmt"

// i is the first pointer and j is the second pointer, one from start and anothe from end
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ { //while swap we need to loop only half the length len(s)/2 else reswap will occur
		s[i], s[j] = s[j], s[i]
		j--
	}
	fmt.Println(s)
}
