package slidingwindow

import "fmt"

func Sliding_window_dynamic_size_longestSumLEK(nums []int, k int) int {
	left, sum, max := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum > k {
			sum -= nums[left]
			left++
		}
		if right-left+1 > max {
			max = right - left + 1
		}
	}
	return max
}

func main_dynamic_size() {
	s := []int{2, 1, 5, 1, 3, 2}
	window_size := 7
	fmt.Println(Sliding_window_dynamic_size_longestSumLEK(s, window_size))
}