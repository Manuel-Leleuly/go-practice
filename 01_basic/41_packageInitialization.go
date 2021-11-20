package main

import (
	"fmt"
	"go-practice/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}