package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchtonouse(group *sync.WaitGroup, number int) {
	defer group.Done()

	group.Add(1)
	fmt.Println("hellow ", number)
	time.Sleep(1 * time.Second)
}
func TestWaaitGroup(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go RunAsynchtonouse(&group, i)
	}
	group.Wait()
	fmt.Println("selesai")
}
