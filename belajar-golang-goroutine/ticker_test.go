package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	isRunning := true

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		isRunning = false
	}()

	for time := range ticker.C {
		fmt.Println(time)
		if !isRunning {
			break
		}
	}
}
