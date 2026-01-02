package slidingwindow

import "fmt"

func Sliding_window_dynamic_size_longestSumLEK(nums []int, K int) int {
	left := 0
	sum := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum > K {
			sum -= nums[left]
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func main_dynamic_size() {
	s := []int{2, 1, 5, 1, 3, 2}
	window_size := 7
	fmt.Println(Sliding_window_dynamic_size_longestSumLEK(s, window_size))
}
