package leetcode

import "log"

func Leet_28(){
	haystack,needle:="abc","c"
	res:=StrStr(haystack,needle)
	log.Println(res)
}

func StrStr(haystack string, needle string) int {
	l:=len(needle)
	for i:=0;i<=len(haystack)-l;i++{
		if haystack[i:l+i]==needle{
			return i
		}
	}
	return -1
}
/*28. Find the Index of the First Occurrence in a String
Easy
Topics
Companies
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

 

Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
 

Constraints:

1 <= haystack.length, needle.length <= 104
haystack and needle consist of only lowercase English characters.
*/