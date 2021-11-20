package belajargolanggoroutines

import (
	"fmt"
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
