package belajargolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

/*
	context with value
*/
func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	/*
		the way you can get the context value is by using Value() method.
		When the value is not in the context you specified, it will go to its parent to get the value.
		When the value is not in its parent, it will go the parent before it and so on until the root parent.
		If the value is not in the root parent, then the method will return <nil>.

		INFO: the Value() method will not search for value in the child or grandchild of the specified context
	*/

	fmt.Println(contextF.Value("f")) // will get the value
	fmt.Println(contextF.Value("c")) // will get the value from its parent
	fmt.Println(contextF.Value("b")) // won't get the value because it doesn't have a parent "b"
	fmt.Println(contextA.Value("b")) // won't get the value because "a" is the root parent
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return // using "break" will only break select and not the for loop
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulate slow process
			}
		}
	}()
	/*
		BEFORE:
			for {
				destination <- counter
				counter++
			}

		The for loop inside the goroutine (which is written in BEFORE) is set to loop forever,
		meaning that the destination channel will always receive the data from counter.

		even though the looping in the TestContextWithCancel stops after n reaches 10,
		because the loop inside the goroutine loops forever,
		the goroutine will always try send a data into the channel (in this case is "destination")
		without any functions or methods written to stop the goroutine from running forever.
		As the result, even though the program stops running, the goroutine will continue its job.
		This is called "goroutine leak"

		AFTER:
			for {
				select{
				case <- ctx.Done():
					return // using "break" will only break select and not the for loop
				default:
					destination <- counter
					counter++
				}
			}
		By sending the context with cancel, the goroutine will stop if the context is done using Done() method
	*/

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel() // sends a cancel signal to context

	time.Sleep(2 * time.Second)

	fmt.Println("Total goroutine", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second) // automates cancel until a certain time period
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total goroutine", runtime.NumGoroutine())
}

/*
	WithDeadline is similar to WithTimeout.
	The difference is that WithDeadline uses real life time as the cancel automation,
	whereas WithTimeout uses a duration set by the user.
*/
func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second)) // automates cancel at a certain time
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total goroutine", runtime.NumGoroutine())
}
