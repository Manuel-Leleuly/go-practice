package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/*
	atomic is another way to use lock and unlock similar to mutex
*/

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x)
}
