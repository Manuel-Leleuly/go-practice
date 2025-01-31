package main

import (
	"fmt"
	"os"
)

func sayHello() {
	fmt.Println("to the rugs topography")
}

func main() {
	sayHello()

	hostName, err := os.Hostname()

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(hostName)
	}
}
