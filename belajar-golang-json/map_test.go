package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{P0001 Apple Mac Book Pro http://example.com/image.png}`
	jsonBytes := []byte(jsonString)

	var result map[string]any
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]any{
		"id":    "P0001",
		"name":  "Apple Mac Book Pro",
		"price": 2000000000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
