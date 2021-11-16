package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Manuel Theodore", "Manuel"))
	fmt.Println(strings.Split("Manuel Theodore", " "))
	fmt.Println(strings.ToLower("Manuel Theodore"))
	fmt.Println(strings.ToUpper("Manuel Theodore"))
	fmt.Println(strings.ToTitle("manuel theodore"))
	fmt.Println(strings.Trim("             Manuel Theodore         ", " "))
	fmt.Println(strings.ReplaceAll("Manuel Manuel Manuel Manuel Manuel", "Manuel", "Theodore"))
}