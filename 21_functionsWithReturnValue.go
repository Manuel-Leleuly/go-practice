package main

import "fmt"

func getHello(name string) string {
	if name == ""{
		return "Hello bro"
	}
	return "Hello " + name
}

func main() {
	result := getHello("Manuel")
	fmt.Println(result)
}