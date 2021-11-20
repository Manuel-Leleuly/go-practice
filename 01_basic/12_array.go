package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Manuel"
	names[1] = "Theodore"
	names[2] = "Leleuly"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	values := [3]int{
		90,
		80,
		95,
	}

	fmt.Println(values)
}