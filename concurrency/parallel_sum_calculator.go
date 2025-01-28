package concurrency

import "fmt"

func sum(slice []int, ch chan int) {
	total := 0
	for _, v := range slice {
		total += v
	}
	ch <- total
}

func Go_routines() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	go sum(numbers[:len(numbers)/2], ch)
	go sum(numbers[len(numbers)/2:], ch)
	x, y := <-ch, <-ch
	fmt.Println("Total Sum: ", x+y,x,y)
}