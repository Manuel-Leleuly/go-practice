package main

import "fmt"

func main() {
	a := 10
	b := 10
	c := a + b
	fmt.Println(c)

	i := 10
	i += 10
	fmt.Println(i)

	i++
	fmt.Println(i)
}