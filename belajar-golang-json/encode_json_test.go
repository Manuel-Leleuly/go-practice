package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	logJson("Manuel")
	logJson(1)
	logJson(true)
	logJson([]string{"Manuel", "Theodore", "Leleuly"})
}
