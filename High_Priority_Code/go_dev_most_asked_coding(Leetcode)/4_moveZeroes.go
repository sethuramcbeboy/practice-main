package godevmostaskedcoding

import (
	"fmt"
)

func moveZeroes(nums []int)  []int{
	ctr :=0
	for i:=0;i<len(nums);i++{
		if nums[i] != 0{
			nums[ctr],nums[i] = nums[i],nums[ctr]   //Bringing Non zero electemts to the front
			ctr++
		}
	}
	return nums
}

func main() {
	nums := []int{1,2,3,0,4,0}
	fmt.Println("Result is ", moveZeroes(nums))
}