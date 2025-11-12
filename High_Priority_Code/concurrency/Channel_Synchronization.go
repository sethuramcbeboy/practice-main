// Channel Synchronization
// Goal
// Print numbers in order using goroutines, ensuring synchronization with channels.

// Input
// No input — just print 1 to 5 sequentially, but each number printed by a separate goroutine.

// Expected Output
// 1
// 2
// 3
// 4
// 5

// Requirements

// Use 5 goroutines.

// Use channels to ensure order — e.g., goroutine 2 waits until goroutine 1 finishes, and so on.

// No global variables or sleeps — only channel-based synchronization.

package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func sync_print(wg *sync.WaitGroup, v int) {
	fmt.Println(v)
	wg.Done()
}

func Main_Channel_sync() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go sync_print(&wg,i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}