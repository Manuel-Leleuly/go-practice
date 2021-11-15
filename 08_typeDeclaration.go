package main

import "fmt"

func main() {
	type NoKTP string

	var noKtpManuel NoKTP = "12312312412452514"
	fmt.Println(noKtpManuel)
}