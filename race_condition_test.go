package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestRacesCondition(t *testing.T) {
	var x = 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("nilai X =", x)
}
