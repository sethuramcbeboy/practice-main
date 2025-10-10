package practice

import "fmt"

// Done using recursion
func fibo_check(n int, sum int) int {
	sum = sum * n
	if n <= 1 {
		return sum
	}
	return fibo_check(n-1, sum)
}

func fibo() {
	n, sum := 5, 1
	res := fibo_check(n, sum)
	fmt.Println(res)
	res2 := fibo_check_loop(n, sum)
	fmt.Println(res2)
}

// 5 -> 120 (5 x 4 x 3 x 2 x 1)

// using loop
func fibo_check_loop(n int, sum int) int {
	for i := n; i > 0; i-- {
		sum *= i
	}
	return sum
}