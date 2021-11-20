package main

import (
	"fmt"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing"{
		return "..."
	}
	return name
}

func main() {
	sayHelloWithFilter("Manuel", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}