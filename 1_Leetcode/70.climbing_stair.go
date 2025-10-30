package leetcode

import "log"

func Leet_70(){
	n:=7
	res:=climbStairs(n)
	log.Println(res)
}

func climbStairs(n int) int {
	if n<4{
		return n
	}
	a:=make([]int,n+1)
	a[1],a[2]=1,2
	for i:=3;i<=n;i++{
		a[i]=a[i-1]+a[i-2]
	}
	return a[n]
}

/*You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

 

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 */