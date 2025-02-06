package leetcode

import "log"

func Leet_67(){
	a,b:="11","1"
	res:=addBinary(a,b)
	log.Println(res)
}

func addBinary(a string, b string) string {
	resp,carry,i,j:="",0,len(a)-1,len(b)-1
	for i>=0 || j>=0 || carry!=0{
		x:=0
		if i>=0{
			x=int(a[i]-'0')									
			i--
		}
		y:=0
		if j>=0{
			y=int(b[j]-'0')
			j--
		}
		sum:=x+y+carry
		resp=string(sum%2+'0') +resp
		carry=sum/2
	}
	return resp
}


/*67. Add Binary
Easy
Topics
Companies
Given two binary strings a and b, return their sum as a binary string.

 

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/