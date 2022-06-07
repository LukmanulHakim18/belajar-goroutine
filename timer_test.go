package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {

	timer := time.NewTimer(5 * time.Second)

	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)

}
func TestTimerAfter(t *testing.T) {

	channel := time.After(5 * time.Second)

	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)

}

func TestTimerAfterFunc(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	time.AfterFunc(5*time.Second, func() {
		defer wg.Done()
		fmt.Println(time.Now())
	})
	fmt.Println(time.Now())
	wg.Wait()
}
