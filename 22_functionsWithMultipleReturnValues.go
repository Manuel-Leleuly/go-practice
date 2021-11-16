package main

import "fmt"

func getFullName() (string, string) {
	return "Manuel", "Leleuly"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}