package concurrency

import (
	"fmt"
	"sync"
)

func worker_with_mu(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup, mu *sync.Mutex) {
	for j := range jobs {
		mu.Lock()
		fmt.Printf("Worker %d processing job %d\n", id, j)
		mu.Unlock()

		results <- j * 2
		wg.Done()
	}
}

func Worker_pool_mu() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup
	var mu sync.Mutex // <-- Mutex added

	for w := 1; w <= 3; w++ {
		go worker_with_mu(w, jobs, results, &wg, &mu)
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
