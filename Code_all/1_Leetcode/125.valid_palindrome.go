package leetcode

import (
	"log"
	"unicode"
)
func Leet_125(){
	s:="A man, a plan, a canal: Panam"
	res:=IsPalindrome(s)
	log.Println(res)
}

func IsPalindrome(s string) bool {
    var a []rune
	for _,v:=range s{
		if unicode.IsLetter(v) || unicode.IsDigit(v){
			a=append(a,unicode.ToLower(v))
		}
	}
	if len(a) == 0 {
		return true
	}
	b:=make([]rune, len(a))
	copy(b,a)
	j:=len(a)-1
	for i:=0;i<len(a)/2;i++{
		b[i],b[j]=b[j],b[i]
		j--
	}
	for i:=range a{
		if a[i]!=b[i]{
			return false
		}
	}
	return true
}


/*125. Valid Palindrome
Easy
Topics
Companies
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

 

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/