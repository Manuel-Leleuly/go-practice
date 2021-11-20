package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	manuel := Man{"Manuel"}
	manuel.Married()
	fmt.Println(manuel)
}