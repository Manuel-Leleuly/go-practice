package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(*address2)

	address4 := new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	address5 := Address{"Subang", "Jawa Barat", ""}
	changeCountryToIndonesia(&address5)
	fmt.Println(address5)
}