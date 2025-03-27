package mapcode

import "fmt"

func countFrequency1(arr []int) map[int]int {
    freq := make(map[int]int)

    for _, num := range arr {
        freq[num]++
    }
    return freq
}

func hashing() {
    arr := []int{1, 2, 2, 3, 1, 1}
    fmt.Println(countFrequency1(arr)) // Output: map[1:3 2:2 3:1]
}
