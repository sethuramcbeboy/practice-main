package concurrency

import (
	"fmt"
	"sync"
)

type Result struct {
	Worker_id int
	v         int
	res       int
}

func Worker_func(id int, ch chan int, ch2 chan Result, wg *sync.WaitGroup) {
	for v := range ch {
		ch2 <- Result{Worker_id: id, v: v, res: v * v}
	}
	wg.Done()
}

func main_worker() {
	var wg sync.WaitGroup
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make([]chan int, 3)
	ch2 := make(chan Result, 10)
	for i := 0; i <= 2; i++ {
		ch[i] = make(chan int, 10)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Worker_func(i, ch[i], ch2, &wg)
	}
	
	a := 0
	for _, v := range s {
		a %= 3
		ch[a] <- v
		a++
	}
	for i := 0; i <= 2; i++ {
		close(ch[i])
	}
	go func() {
		wg.Wait()
		close(ch2)
	}()
	for v := range ch2 {
		fmt.Printf("Worker %d done the job %d and its result is %d \n", v.Worker_id, v.v, v.res)
	}
}
