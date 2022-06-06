package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var wg = sync.WaitGroup{}

func WaitCondition(value int) {
	defer wg.Done()

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go WaitCondition(i)

	}
	// mengirim  signal satu persatu
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("kirim signal")
			cond.Signal()
		}
	}()

	// mengirim signal brodcast sekaligus
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("kirim brodcast")
	// 	cond.Broadcast()
	// }()
	wg.Wait()
	fmt.Println("selesai")
}
