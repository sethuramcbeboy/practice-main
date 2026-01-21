// 1. Spawn 3 threads, each representing a different camera:
// 2. Every 100–200 ms, each thread pushes a "frame" into a shared buffer:
// e.g. `Frame{camera_id: "Camera1", timestamp: 1717938123456}`
// 3. Another thread continuously pops frames and logs them like:
// e.g. `[Camera2] Frame @ 1717938123456
//
//	[Camera1] Frame @ 1717938123488`
//
// 4. Implement your own thread-safe queue using mutex for the producer-consumer setup

package tradelab

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)
////////////////////////////////////////////////////////////////
// simple one
func cam(a string, ch chan int) {
	for v := range ch {
		fmt.Printf("[%s] Frame @ %d \n", a, v)
	}
}

func mainM() {
	ch := make(chan int, 100)
	for i := 0; i <= 2; i++ {
		a := "camera" + strconv.Itoa(i)
		//fmt.Println(a)
		go cam(a, ch)
	}
	for {
		ch <- int(time.Now().UnixMicro())   // time stamp generating
		time.Sleep(250 * time.Millisecond)
	}
}
////////////////////////////////////////////////////////////////
//Advanced one

type Frame struct {
	id string
	ts int64
}

type Queue struct {
	m sync.Mutex
	c *sync.Cond
	q []Frame
}

func NewQueue() *Queue {
	q := &Queue{}
	q.c = sync.NewCond(&q.m)
	return q
}

func (q *Queue) Push(f Frame) {
	q.m.Lock()
	q.q = append(q.q, f)
	q.m.Unlock()
	q.c.Signal() // notify waiting goroutine that a new frame is available
}

func (q *Queue) Pop() Frame {
	q.m.Lock()
	for len(q.q) == 0 {
		q.c.Wait() // waits until Push signals that data is available
	}
	f := q.q[0]
	q.q = q.q[1:]
	q.m.Unlock()
	return f
}

func camera(id string, q *Queue) {
	for {
		time.Sleep(time.Duration(rand.Intn(100)+100) * time.Millisecond) // for the time delay of 100 - 200 ms
		q.Push(Frame{id, time.Now().UnixMilli()})
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	q := NewQueue()

	// 3 producer goroutines (cameras)
	var s []string
	s = []string{"c1", "c2", "c3"}

	for i := 0; i <= 2; i++ {
		go camera(s[i], q)
	}
	// single consumer goroutine
	for {
		f := q.Pop()
		// FIX: fmt.Println doesn’t support format specifiers — use fmt.Printf
		fmt.Printf("[%s] Frame @ %d\n", f.id, f.ts)
	}
}
