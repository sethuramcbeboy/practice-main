package godevmostaskedcoding

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for _, v := range m {
		if v > 1 {
			return true
		}
	}
	return false
}

func main_dup() {
	nums := []int{1, 2, 3, 4, 2}
	fmt.Println("Result is ", containsDuplicate(nums))
}
