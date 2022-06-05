package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalence struct {
	sync.Mutex
	Name    string
	Balance int
}

func (ub *UserBalence) Lock() {
	ub.Mutex.Lock()
}
func (ub *UserBalence) Unlock() {
	ub.Mutex.Unlock()
}

func (ub *UserBalence) Change(amount int) {
	ub.Balance = ub.Balance + amount
}

func Transfer(from *UserBalence, to *UserBalence, amount int) {
	from.Lock()
	fmt.Println("locking user balence from ", from.Name)
	from.Change(-amount)

	time.Sleep(1 * time.Second)

	to.Lock()
	fmt.Println("locking user balence to ", to.Name)
	to.Change(amount)

	time.Sleep(1 * time.Second)

	from.Unlock()
	to.Unlock()

}

func TestDeadlock(t *testing.T) {
	user1 := UserBalence{
		Name:    "Lukman",
		Balance: 1000000,
	}
	user2 := UserBalence{
		Name:    "Hakim",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(3 * time.Second)

	fmt.Println("User ", user1.Name, ", Balence ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balence ", user2.Balance)
}
