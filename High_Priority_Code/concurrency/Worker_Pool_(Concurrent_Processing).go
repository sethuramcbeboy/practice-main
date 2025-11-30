package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func worker_func(wg *sync.WaitGroup, id int, job chan int, result chan int) {
	for v := range job {
		fmt.Printf("Worker %d doing job %d \n", id, v)
		time.Sleep(1 * time.Second)
		result <- v
	}
	wg.Done()
}

func Main_worker_pool() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	job := make(chan int, 10)
	result := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i <= 2; i++ {
		wg.Add(1)
		go worker_func(&wg, i, job, result)
	}
	for _, v := range a {
		job <- v
	}
	close(job)

	go func() {
		wg.Wait()
		close(result)
	}()

	for v := range result {
		fmt.Println("result =", v)
	}
}
