package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {

	pool := sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}
	wg := sync.WaitGroup{}
	pool.Put("luke")
	pool.Put("man")
	pool.Put("nul")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := pool.Get()
			fmt.Println(i, data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
	wg.Wait()
	fmt.Println("selesai")
}
