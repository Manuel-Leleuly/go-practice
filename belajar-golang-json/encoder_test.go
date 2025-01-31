package belajargolangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Manuel",
		MiddleName: "Theodore",
		LastName:   "Leleuly",
	}
	_ = encoder.Encode(customer)
}
