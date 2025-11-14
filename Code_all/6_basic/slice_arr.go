package practice

import "fmt"

func Slice_arr_operation(req [10]int) (res []int) {
	var s []int
	s = req[:5]
	fmt.Println(s)
	for i := 5; i < len(req); i++ {
		s = append(s, req[i])
	}
	return s
}
