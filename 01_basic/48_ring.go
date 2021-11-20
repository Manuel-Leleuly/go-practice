package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)
	// data.Value = "Manuel"
	// data2 := data.Next()
	// data2.Value = "Theodore"

	for i := 0; i < data.Len(); i++{
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(i interface{}) {
		fmt.Println(i)
	})
}