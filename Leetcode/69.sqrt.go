package leetcode

import "log"

func Leet_69(){
	x:=50
	res:=MySqrt(x)
	log.Println(res)
}

func MySqrt(x int) int {
    if x==0 || x==1{
		return x
	}
	res:=1
	for res*res<=x{
		res++
	}
	return res-1
}

/*69. Sqrt(x)
Solved
Easy
Topics
Companies
Hint
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
 

Example 1:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.
Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
*/