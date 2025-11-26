package godevmostaskedcoding

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string{
	sarr := strings.Fields(s)   // wont consider space in front and back only take the string separated by gap and return []string
	j:=len(sarr)-1
	for i:=0;i<len(sarr)/2;i++{  // divide by 2 else it will revert the process happend
		sarr[i],sarr[j] = sarr[j],sarr[i]
		j--
	}
	return strings.Join(sarr," ")
}

func main() {
	s := "  hello world  "
	fmt.Println(reverseWords(s))
}


//high load taking one
func reverseWords_highload(s string) string {
	ctr, m := 0, make(map[int]string)
	str := ""
	for _, v := range s {
		if v == ' ' {
			if str != "" {
				m[ctr] = str
				ctr += 1
			}
			str = ""
		} else {
			str += string(v)
		}
	}
	if str != "" {
		m[ctr] = str
	}
	res_Str := ""
	for i := ctr; i >= 0; i-- {
		res_Str += m[i]
		if i > 0 && m[i] != "" {
			res_Str += " "
		}
	}
	return res_Str
}