package tradelab

import (
	"fmt"
	"sync"
	"time"
)

// func Writefunc(s []string) {
// 	for _, v := range s {
// 		fmt.Println(v)
// 	}
// }

// func main() {
//     s := []string{"a", "b", "c", "d"}
//     go Writefunc(s)
// }

// output:- nothing is printed, because  once the Writefunc called with Go_routines,the main function ends. so that the main thread will close
// which causes the go routine also closed.

// solution:-
// 	easy:- time.Sleep(150 * time.Millisecond)
// 	advanced :- wait group   -> Attracts interviewer

// easy:- time.Sleep(150 * time.Millisecond)
func Writefunc(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Main() {
	s := []string{"a", "b", "c", "d"}
	go Writefunc(s)
	time.Sleep(1 * time.Minute)
}

//////////////////////////////////////////////

// advanced :- wait group   -> Attracts interviewer
func Writefunc_waitgroup(s []string, wg *sync.WaitGroup) {
	for _, v := range s {
		fmt.Println(v)
	}
	wg.Done()
}

func Main_wait_group() {
	s := []string{"A", "B", "C", "D"}
	var wg sync.WaitGroup
	wg.Add(1)
	go Writefunc_waitgroup(s, &wg)
	wg.Wait()
}

// using channel (Very very simple but highly attract interviewer)
////////////////////////////////////////////////
func Print_func(s []string, ch chan bool) {
	fmt.Println(s)
	ch <- true
}

func main_channel() {
	a := []string{"A", "B", "C", "D"}
	ch := make(chan bool)
	go Print_func(a, ch)
	<-ch
}