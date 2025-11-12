package concurrency

import (
	"log"
	"sync"
)

func con_sum(wg *sync.WaitGroup,a int,v int,ch chan int){
	defer wg.Done()
	log.Println(a) // to track all the go routie worked
	ch <- v*v
}

func Main_con_sum(){
	a := []int{2,4,6,8,10}
	ch := make(chan int,10)
	var wg sync.WaitGroup

	for i,v := range a{
		wg.Add(1)
		go con_sum(&wg,i,v,ch)
	}
	go func ()  {
		wg.Wait()
		close(ch)
	}()

	sum := 0
	for v := range ch{
		sum += v
	}
	log.Println("Sum of Square = ",sum)
}