package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestEncode(t *testing.T) {
	LogJson("Manuel")
	LogJson(1)
	LogJson(true)
	LogJson([]string{"Manuel", "Theodore", "Leleuly"})
}

/*
	Because json.Marshal takes any data with data type of interface{} (meaning that all type of data is allowed),
	we can encode any data as we want (string, number, boolean, etc.).
	Even thought it's allowed in the function, it's not allowed in the JSON contract.
	If we want to follow json.org, the JSON data has to be an object or an array.
	Then the value inside of it can be anything we want.
*/
