package main

import "fmt"

func main() {
	name := "Manuel"

	if name == "Manuel" {
		fmt.Println("Hello Manuel")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	if length := len(name); length > 5{
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}