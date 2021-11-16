package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your host")
	user := flag.String("user", "root", "Put your database user")
	pass := flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Pass:", *pass)
}