package leetcode

import (
	"fmt"
)

func RemoveDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	j:=0
	for i:=1;i<len(nums);i++{
		if nums[i]!=nums[j]{
			j++
			nums[j]=nums[i]
		}
    }
    return nums[:j+1]
}

func Leet_26() {
	nums := []int{1,1,2,3,4,4,5,6,7,8}
	res := RemoveDuplicates(nums)
	fmt.Println("res = ",res) // Expected Output: [1, 2, 3, 4]
}
