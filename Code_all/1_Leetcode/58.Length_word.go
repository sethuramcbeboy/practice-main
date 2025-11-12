package leetcode

import "log"

func Leet_58(){
    s := "   fly me   to   the moon  "
	res:=lengthOfLastWord(s)
	log.Println(res)
}
func lengthOfLastWord(s string) int {
	if len(s)==0{
		return 0
	}
	ctr:=0
	for i:=len(s)-1;i>=0;i--{
		if s[i]!=' '{
			ctr++
		}else{
			if ctr>0{
				break
			}
		}
	}
	return ctr
}

/*58. Length of Last Word
Easy
Topics
Companies
Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal 
substring
 consisting of non-space characters only.

 

Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.
*/