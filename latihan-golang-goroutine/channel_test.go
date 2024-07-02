package latihan_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSelectChannelDefault(t *testing.T) {
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
			fmt.Println("data channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data channel 2", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}

func TestSelectChannel(t *testing.T) {
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
			fmt.Println("data channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Request input ke " + strconv.Itoa(i)
		}
		close(channel) //jika tidak diclose, process akan deadlock
	}()

	for data := range channel {
		fmt.Println("Menerima data ", data)
	}
	fmt.Println("selesai")
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	go func() {
		channel <- "aant can"
		channel <- "aant sultan"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "rahmanya"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("Output", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string) // seperti future api dalam Java
	defer close(channel)
	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(5 * time.Second)
}
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "elaina"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string) // seperti future api dalam Java
	defer close(channel)
	//channel <- "aant" // mengirim data ke channel
	go GiveMeResponse(channel)
	data := <-channel // menerima data

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestChannel(t *testing.T) {
	channel := make(chan string) // seperti future api dalam Java
	defer close(channel)
	//channel <- "aant" // mengirim data ke channel
	//

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "aant"
		fmt.Println("Selesai")
	}()

	data := <-channel // menerima data

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
