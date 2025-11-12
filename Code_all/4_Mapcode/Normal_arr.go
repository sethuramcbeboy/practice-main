package mapcode

import (
	"fmt"
	"log"

)

func countFrequency(arr []int) map[int]int {
	freq := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		freq[arr[i]] = count
        log.Println(arr[i],"=====",freq)
	}
	return freq
}

func Map_hash1() {
	arr := []int{1, 2, 2, 3, 1, 1}
	fmt.Println(countFrequency(arr)) // Output: map[1:3 2:2 3:1]
}
