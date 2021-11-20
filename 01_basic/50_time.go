package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	
	fmt.Println(now)
	fmt.Println(now.Day())
	fmt.Println(now.Month())
	fmt.Println(now.Year())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2021, 11,16, 20, 16, 20,2000, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2021-11-16")
	fmt.Println(parse)
}