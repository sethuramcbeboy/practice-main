package practice

import "fmt"

func Merge() {
	var s []int
	var l []int
	s = append(s, 1, 2, 3, 4, 5)
	l = append(l, 6, 7, 8, 9, 10)
	// for _, v := range l {
	// 	s = append(s, v)
	// }
	//         or
	s=append(s, l...) 
	fmt.Println(s)
}
