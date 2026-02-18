package twopointers

import "fmt"

func maxArea(height []int) int {
	max := 0
	i, j := 0, len(height)-1
	for i < j {
		h := min(height[i], height[j])
		d := j - i
		if h*d > max {
			max = h * d
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main_m() {
	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(s))
}
