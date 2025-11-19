package practice

import (
	"fmt"
	"strconv"
	"strings"
)

func Comma_string_separate(s string) []int {
	value := strings.Split(s, ",")
	length := len(value)
	var err error
	slice := make([]int, length)
	for i, v := range value {
		a := strings.Trim(v, " ")
		slice[i], err = strconv.Atoi(a)
		if err !=nil{
			fmt.Println(err)
		}
	}
	return slice
}	
