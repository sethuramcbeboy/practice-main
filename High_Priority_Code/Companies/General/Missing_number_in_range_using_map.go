package general

import "fmt"

// given list containig numbers 1 to N with missing number, find the missing numbers
// a := [1, 2, 4, 5, 6, 10, 16, 18] N = 10
// output :- [3,7,8,9] -> under 10 missing value as N = 10
func Find_missing_number_in_range(s []int, n int) []int {
	var res []int
	m := make(map[int]int)
	for _, v := range s {
		m[v]++
	}
	for i := 1; i <= n; i++ {
		if m[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}

func main2() {
	s := []int{1, 2, 4, 5, 6, 10, 16, 18}
	N := 10
	fmt.Println(Find_missing_number_in_range(s, N))
}
