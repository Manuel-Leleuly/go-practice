package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("customer.json")
	decoder := json.NewDecoder(reader)

	customer := Customer{}
	_ = decoder.Decode(&customer)

	fmt.Println(customer)
}
