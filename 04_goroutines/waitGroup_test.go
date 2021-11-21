package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAynchronous(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}
