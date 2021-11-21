package belajargolanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

/*
	REMINDER: goroutines work concurrently instead of parallel.
*/

func TestCreateGoRoutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second) // my laptop is slow, alright :"(
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Manuel Theodore Leleuly"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

/*
	You can set the channel's behaviour to either send or receive data

	for example:
	chan<- string; it means that the channel can only be used to receive data from a goroutine
	<-chan string: it means that the channel can only be used to send data from a goroutine
*/

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Manuel Theodore Leleuly"
}

func OnlyOut(channel <-chan string) {
	data := <-channel

	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // determine the number of buffer given to the channel
	defer close(channel)

	go func() {
		channel <- "Manuel"
		channel <- "Theodore"
		channel <- "Leleuly"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

/*
	If you are not sure how many times the channel will receive a data, you can use a ranged channel using for loop
*/

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
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
			fmt.Println("Data dari Channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2:", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}
