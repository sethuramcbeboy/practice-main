package godevmostaskedcoding

import "fmt"


func twoSum(nums []int, target int) []int {
    for i:=0;i<len(nums)-1;i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i]+nums[j] == target{
                res := []int{i,j}
                return res
            }
        }
    }
    return nil
}

func main_two_sum(){
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println("Result is ",twoSum(nums,target))
}