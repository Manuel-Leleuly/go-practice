package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (c Customer) hasBirthday() {
	c.Age++
}

func main() {
	var manuel Customer
	manuel.Name = "Manuel"
	manuel.Address = "Bandung"
	manuel.Age = 15

	fmt.Println(manuel)

	theodore := Customer{
		Name: "Theodore",
		Address: "Jakarta",
		Age: 16,
	}
	fmt.Println(theodore)

	leleuly := Customer{"Leleuly", "Indonesia", 17}
	fmt.Println(leleuly)
}