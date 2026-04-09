package comcast

import (
	"fmt"
	"sync"
)

// Write a program that should process a series of jobs using fixed number of go routines, 30 jobs using 5 goroutines

type Result struct {
	Id   int
	Jobs int
}

func Worker(I int, ch chan int, ch2 chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		ch2 <- Result{Id: I, Jobs: v}
	}
}

func main_() {
	ch := make([]chan int, 5)
	for i := 0; i < 5; i++ {
		ch[i] = make(chan int, 6)
	}
	ch2 := make(chan Result, 10)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Worker(i, ch[i], ch2, &wg)
	}
	for i := 0; i < 30; i++ {
		ch[i%5] <- i
	}
	for i := 0; i < 5; i++ {
		close(ch[i])
	}
	go func() {
		wg.Wait()
		close(ch2)
	}()
	for v := range ch2 {
		fmt.Printf("Worker %d had done the job %d \n", v.Id, v.Jobs)
	}
}