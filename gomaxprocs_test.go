package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(3 * time.Second)
		}()
	}
	totalCpu := runtime.NumCPU()

	fmt.Println("total Cpu:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total Thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total Goroutine:", totalGoroutine)
	wg.Wait()
}
func TestChangeThreadNumber(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(3 * time.Second)
		}()
	}
	totalCpu := runtime.NumCPU()

	fmt.Println("total Cpu:", totalCpu)
	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total Thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total Goroutine:", totalGoroutine)
	wg.Wait()
}
