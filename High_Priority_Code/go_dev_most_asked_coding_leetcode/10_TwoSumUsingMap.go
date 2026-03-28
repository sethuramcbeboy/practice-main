package godevmostaskedcoding

import "fmt"

func twoSumUsingMap(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if idx, found := m[complement]; found {
            return []int{idx, i}
        }
        m[num] = i
    }
    return []int{}
}

func main1() {
    a := []int{2, 7, 9, 10}
    t := 12
    fmt.Println(twoSumUsingMap(a, t))
}