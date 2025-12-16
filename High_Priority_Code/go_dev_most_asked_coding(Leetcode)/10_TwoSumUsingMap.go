package godevmostaskedcoding

import "fmt"

func twoSumUsingMap(nums []int, target int) []int {
    m := make(map[int]int)

    for i, num := range nums {
        needed := target - num

        if m[needed] != 0 {
            return []int{m[needed] - 1, i}
        }

        // store index + 1
        m[num] = i + 1
        fmt.Println(m)
    }
    return nil
}

func main() {
    a := []int{2, 7, 9, 10}
    t := 12
    fmt.Println(twoSumUsingMap(a, t))
}