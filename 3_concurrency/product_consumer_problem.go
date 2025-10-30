package concurrency

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Produced:", i)
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer(ch chan int) {
	for num := range ch {
		fmt.Println("Consumed:", num)
		time.Sleep(2 * time.Second)
	}
}

func Gorountinne_channel() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
	time.Sleep(20 * time.Second)
}


// Explanation:- first function generate data which is sent to second function using channel

// //using bufferd channel
// func producer(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		ch <- i // Send data into the channel
// 		fmt.Println("Produced:", i)
// 		time.Sleep(time.Second) // Simulate some work
// 	}
// 	close(ch) // Close the channel after producing all data
// }

// func consumer(ch chan int) {
// 	for num := range ch { // Continuously receive data from the channel
// 		fmt.Println("Consumed:", num)
// 		time.Sleep(2 * time.Second) // Simulate some work
// 	}
// }

// func Gorountinne_channel() {
// 	ch := make(chan int, 5) // Create a buffered channel with capacity 5
// 	go producer(ch)         // Start the producer goroutine
// 	go consumer(ch)         // Start the consumer goroutine
// 	time.Sleep(20 * time.Second) // Allow goroutines to run for 20 seconds
// }