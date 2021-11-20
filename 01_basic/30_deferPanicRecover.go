package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int){
	defer logging()
	fmt.Println("Run application")
	result := 10/ value
	fmt.Println("Result", result) 
}

func endApp(){
	message := recover()
	if message != nil{
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	// runApplication(0)
	runApp(true)
}