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
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hoobies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {
	bytes, _ := json.Marshal(Customer{
		FirstName:  "Manuel",
		MiddleName: "Theodore",
		LastName:   "Leleuly",
	})
	fmt.Println(string(bytes))
}
