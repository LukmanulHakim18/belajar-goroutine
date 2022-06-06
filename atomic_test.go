package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("nilai X =", x)
}
