package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Manuel"
	middleName = "Theodore"
	lastName = "Leleuly"
	return
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}