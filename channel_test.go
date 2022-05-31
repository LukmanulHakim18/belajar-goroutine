package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Lukman"
	}()

	data := <-channel
	fmt.Println("ini isi chanel", data)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Lukmanul Hakim"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println("ini isi chanel", data)
	time.Sleep(3 * time.Second)
}

func OnlyIN(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "SecreateData"
}
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("ini data :", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyOut(channel)
	go OnlyIN(channel)

	time.Sleep(5 * time.Second)

}

// Buffered Channel
// digunakan untuk penyimpanan berdasarkan antrian chanel
// tempat menunggu chanel sebelum masuk ke dalam goroutine

func TestBufferdChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)
	go func() {

		channel <- "Anbukestra"
		channel <- "Anbukestra1"
		channel <- "Anbukestra2"

	}()
	go func() {

		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println("selesai")
}
