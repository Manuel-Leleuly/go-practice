package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street, Country, PostalCode string
}

type Customer struct {
	FirstName, MiddleName, LastName string
	Age                             int
	Married                         bool
	Hobbies                         []string
	Addresses                       []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Manuel",
		MiddleName: "Theodore",
		LastName:   "Leleuly",
		Age:        26,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
