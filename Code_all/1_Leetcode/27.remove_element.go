package leetcode

import "log"

func Leet_27(){
	nums:=[]int{1,2,3,4,5,6,7,8,3,2,3}
	val:=3
	res:=nums[:RemoveElement(nums,val)]
	log.Println(res)
}

func RemoveElement(nums []int,val int) int{
	 j:=0
	 for i:=0;i<len(nums);i++{
		if nums[i]!=val{
			nums[j]=nums[i]
			j++
		}
	 }
	 return j
}