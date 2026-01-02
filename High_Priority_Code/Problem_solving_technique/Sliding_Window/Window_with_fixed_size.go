package slidingwindow

import "fmt"

func Sliding_window_with_fixed_size(s []int, k int) int {
	sum, max := 0, 0
	for i := 0; i < k; i++ {
		sum += s[i]
	}
	for i := 1; i < len(s); i++ {
		if k+i-1 <= len(s)-1 {
			sum = sum - s[i-1] + s[k+i-1]  // in sum removing old element and add new 
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

// More optimal
func SlidingWindowFixed(s []int, k int) int {
	if len(s) < k {
		return 0
	}
	sum, max := 0, 0

	// Step 1: sum of first window
	for i := 0; i < k; i++ {
		sum += s[i]
	}

	// Step 2: slide the window
	for i := k; i < len(s); i++ {
		sum += s[i]   // Add new element
		sum -= s[i-k] // Remove firt element val in cuttent sum
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	window_size := 3
	fmt.Println("From the sub array max is ", Sliding_window_with_fixed_size(s, window_size))
}
