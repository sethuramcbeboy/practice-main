package leetcode

import (
	"log"
)

func Leet_35(){
	nums,target:=[]int{1,3,5,6},7
	res:=searchInsert(nums,target)
	log.Println(res)
}

func searchInsert(nums []int, target int) int {
	for i,_:=range nums{
		if nums[i]==target{
			return i
		}
	}
	for i:=0;i<len(nums);i++{
		if target<nums[i]{
			return i
		}
	}
	return len(nums)
}

/*35. Search Insert Position
Easy
Topics
Companies
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, 
return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

 

Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4
*/