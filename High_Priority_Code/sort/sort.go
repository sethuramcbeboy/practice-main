package sort

import (
	"fmt"
)

/* ---- 1️⃣ Bubble Sort ----
🔹 Logic: Repeatedly swap adjacent elements if they’re in the wrong order.
🔹 Best for: Easy understanding, small inputs.
🔹 Time: O(n²)
*/
func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

/* ---- 2️⃣ Selection Sort ----
🔹 Logic: Find the smallest element in the unsorted part and place it at the beginning.
🔹 Time: O(n²), fewer swaps than bubble sort.
*/
func selectionSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

/* ---- 3️⃣ Insertion Sort ----
🔹 Logic: Builds the sorted list one element at a time by inserting in the correct position.
🔹 Time: O(n²), but faster than bubble for nearly sorted data.
*/
func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}

/* ---- 4️⃣ Quick Sort ----
🔹 Logic: Pick a pivot, partition elements, and recursively sort sublists.
🔹 Time: O(n log n) average, O(n²) worst-case.
*/
func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	pivot := a[len(a)/2]
	left, right := []int{}, []int{}
	for _, v := range a {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

/* ---- 5️⃣ Merge Sort ----
🔹 Logic: Divide into halves, sort each, and merge sorted halves.
🔹 Time: O(n log n), stable sorting.
*/
func mergeSort(arr []int) []int {
	// Base case: already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Step 1: split
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Step 2: merge
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	// Compare and merge
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	// Add remaining elements
	result = append(result, left...)
	result = append(result, right...)

	return result
}

// ---- 🧪 MAIN ----
func main() {
	data := []int{5, 3, 8, 4, 2}
	fmt.Println("Original:", data)

	a1 := append([]int{}, data...)
	bubbleSort(a1)
	fmt.Println("Bubble Sort:   ", a1)

	a2 := append([]int{}, data...)
	selectionSort(a2)
	fmt.Println("Selection Sort:", a2)

	a3 := append([]int{}, data...)
	insertionSort(a3)
	fmt.Println("Insertion Sort:", a3)

	a4 := append([]int{}, data...)
	fmt.Println("Quick Sort:    ", quickSort(a4))

	a5 := append([]int{}, data...)
	fmt.Println("Merge Sort:    ", mergeSort(a5))
}


/*
Original: [5 3 8 4 2]
Bubble Sort:    [2 3 4 5 8]
Selection Sort: [2 3 4 5 8]
Insertion Sort: [2 3 4 5 8]
Quick Sort:     [2 3 4 5 8]
Merge Sort:     [2 3 4 5 8]
*/