package leetcode

import (
	"fmt"
)

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
    j:=0
    for i:=1;i<len(nums);i++{
        if nums[j]!=nums[i]{
            nums[j]=nums[i]
            j++
        }   
    }
    return j+1
}

func Leet_26() {
	nums := []int{1, 1, 2, 3, 4,4,5,6,7,8}
	newLength := RemoveDuplicates(nums)
	res := nums[:newLength]
	fmt.Println(res) // Expected Output: [1, 2, 3, 4]
}
