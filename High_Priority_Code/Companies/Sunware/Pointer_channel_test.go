package sunware

import "fmt"

// ///////////////////////////////////////////////////
func update(x *int) {
	*x = 10
}

func main_() {
	a := 5
	update(&a)
	fmt.Println(a)
}

// Output = 10

/////////////////////////////////////////////////////

func main_c() {
	ch := make(chan int)

	ch <- 5 // It won't goes to next line
	fmt.Println(<-ch)
}

// output = dead lock

// ch <- 5
//blocks because there is no goroutine receiving from the channel at that moment.
//Since both send and receive are in the same goroutine, the program waits forever and causes a deadlock.
//////////////////////////////////////////////////////
