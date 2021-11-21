package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

/*
	sync.Once{} is used to make sure that the function that you want to run only runs once.
	If there are more than one goroutines, the first goroutine will excecute the function,
	whereas the rest of them will be ignored.
*/
func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
