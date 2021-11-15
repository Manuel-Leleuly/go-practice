package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // int32 is bigger than int8 resulting in the value resets to the lowest value and continue reseting until the value from int32 is finished

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Manuel"
	var m = name[0]
	var mString = string(m)
	fmt.Println(mString)
}