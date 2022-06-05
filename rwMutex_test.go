package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balence int
}

func (account *BankAccount) AddBalence(amount int) {
	account.RWMutex.Lock()
	account.Balence = account.Balence + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalence() int {
	account.RWMutex.RLock()
	balence := account.Balence
	account.RWMutex.RUnlock()
	return balence
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalence(1)
				fmt.Println(account.GetBalence())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(account.GetBalence())
}
