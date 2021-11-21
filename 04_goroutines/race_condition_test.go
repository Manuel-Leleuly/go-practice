package belajargolanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

/*
	when using goroutine, even though it runs concurrently, it can also run parallelly (I don't know if that's a word).
	Mainly because there are couple threads that can be run parallel to each other.
	Therefore, there will be an issue when we modifiy the same variable using more than one goroutine.
	This is called Race Condition.
*/

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i <= 100; i++ {
		go func() {
			for j := 0; j <= 100; j++ {
				x++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x)
}
