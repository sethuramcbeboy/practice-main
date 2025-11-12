package concurrency

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		results <- j * 2
		wg.Done()
	}
}

func Worker_pool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		wg.Add(1)
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("Result:", result)
	}
}
