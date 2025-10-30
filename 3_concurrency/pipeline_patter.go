package concurrency

import "fmt"

func stage1(in chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * 2
		}
		close(out)
	}()
	return out
}

func stage2(in chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i + 1
		}
		close(out)
	}()
	return out
}

func Pipeline_pattern() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	out1 := stage1(ch)
	out2 := stage2(out1)   //out1 is used as request for out2

	for result := range out2 {
		fmt.Println("Result:", result)
	}
}


/*
A ch channel is initialized to supply input data.
A goroutine is launched to send numbers (0 to 9) into the ch channel, then close it.
The stage1 function processes the input from ch (multiplies by 2).
The stage2 function further processes the output of stage1 (adds 1).
The main goroutine reads the final results from the out2 channel and prints them.
*/