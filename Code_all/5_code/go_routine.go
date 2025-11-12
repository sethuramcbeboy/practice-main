package practice

import (
	"fmt"
	"time"
)

func Goroutine(s string){
	go sample1(s)
	time.Sleep(1 * time.Millisecond)
}

func sample1(s string){
	fmt.Println(s)
}