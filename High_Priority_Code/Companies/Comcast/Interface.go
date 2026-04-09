package comcast

import "fmt"

// interface based task runner s := []int{10,20,30} pass through series of functions, add 1 to each , multiply by 2 and square and print output, out put - 484 1764 3844

type Task interface {
	Execute(int) int
}

type Add struct{}
type Multiply struct{}
type Square struct{}

func (a Add) Execute(v int) int {
	return v + 1
}

func (m Multiply) Execute(v int) int {
	return v * 2
}

func (s Square) Execute(v int) int {
	return v * v
}

func Run(t Task, ch chan int) chan int{
	res := make(chan int)
	go func ()  {
		defer close(res)
		for v := range ch{
			res <- t.Execute(v)
		}	
	}()
	return res
}

func main() {
	arr := []int{10,20,30}
	ch := make(chan int)
	add := Run(Add{},ch)
	mul := Run(Multiply{},add)
	squ := Run(Square{},mul)
	go func ()  {
		for _,v := range arr{
			ch <- v
		}
		close(ch)
	}()
	for v := range squ{
		fmt.Println(v)
	}
}