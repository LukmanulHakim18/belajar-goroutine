package goroutines

import (
	"fmt"
	"strconv"
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
	time.Sleep(1 * time.Second)
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
		time.Sleep(2 * time.Second)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("selesai")
}

// Range  Channel
// digunakan untuk mengambil data dari channel yang jumlahnya tidak di ketahui

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("selesai")
}

// select channel
// untuk mengambil data dari seluruh data channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	close(channel1)
	close(channel2)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		default:
			fmt.Println("menunggu data dari channel")
		}
		if counter == 2 {
			break
		}
	}
}
