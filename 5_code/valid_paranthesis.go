package practice

import "log"

// func Valid_Paranthesis_1(s string) bool {
// 	p1,p2,c1,c2,s1,s2,l:=0,0,0,0,0,0,len(s)
// 	for i:=0;i<l;i++{
// 	   switch s[i]{
// 		   case '(':
// 			   p1+=1
// 		   case ')':
// 			   p2+=1
// 		   case '[':
// 			   s1+=1
// 		   case ']':
// 			   s2+=1
// 		   case '{':
// 			   c1+=1
// 		   case '}':
// 			   c2+=1
// 	   }
// 	}
// 	if p1==p2 && s1==s2 && c1==c2{
// 	   return true
// 	}
// 	return false
// }                            //It will not able to check proper closing is there or not,{{()}} valid but {({)}} invalid, but this code cannot ensure it.

func Valid_Paranthesis(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {     //'' - literals or ascii value it take , "" - string
			log.Println("====",string(char))
			stack = append(stack, char)
		} else if closing, ok := pairs[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != closing {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

//using byte data type 
// Instead of range in for loop, we need to use index while looping for byte
// rune will store as unit64 while byte will store ascii character

func Valid_Paranthesis_byte(s string) bool {
	stack := []byte{}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
        char := s[i]

		if char == '(' || char == '[' || char == '{' {     //'' - literals or ascii value it take , "" - string
			log.Println("====",string(char))
			stack = append(stack, char)
		} else if closing, ok := pairs[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != closing {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}