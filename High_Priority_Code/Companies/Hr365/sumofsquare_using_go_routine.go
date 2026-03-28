package hr365

import (
	"fmt"
	"sync"
)

// Write a go program that calculate the sum of squares of numbers from 1 to N.
// Each go routine should be responsible for calculating the square of a number
// and the resule should be sent back using channel to calculate the final sum.
// Input N = 10
func Square_worker(v int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- v * v
}

func main1() {
	sum := 0
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go Square_worker(i, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for v := range ch {
		sum += v
	}
	fmt.Println(sum)
}
