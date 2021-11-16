package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, ". My name is", customer.Name)
}

func main() {
	var manuel Customer
	manuel.Name = "Manuel"
	manuel.Address = "Bandung"
	manuel.Age = 15

	manuel.sayHello("Joko")

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