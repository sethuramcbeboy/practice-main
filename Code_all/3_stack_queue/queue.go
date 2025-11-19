// only difference is at removing for stack s = s[:len(s)-1] where in queueu we removing first element q = q[1:]
package stackqueue

import "fmt"

type queue struct {
	v []int
}

func enqueue(a int,q *queue){
	q.v = append(q.v, a)
	fmt.Println(q.v)
}

func dequeue(q *queue){
	if len(q.v) <= 0 {
		return
	}
	q.v = q.v[1:]
	fmt.Println(q.v)
}

func main_queue() {
	q := &queue{}
	enqueue(1,q)
	enqueue(2,q)
	enqueue(3,q)
	enqueue(4,q)
	dequeue(q)
	dequeue(q)
	dequeue(q)
	dequeue(q)
}
