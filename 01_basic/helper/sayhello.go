package helper

import "fmt"

// use uppercase for the function to be exported
func SayHello(name string) {
	fmt.Println("Hello", name)
}