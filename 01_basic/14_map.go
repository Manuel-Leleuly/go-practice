package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Manuel",
		"address": "Indonesia",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Manuel"
	book["ups"] = "salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}