package main

import "fmt"

func main() {
	ujian := 80
	absensi := 80

	lulusUjian := ujian >= 80
	lulusAbsensi := absensi >= 80

	lulus := lulusAbsensi && lulusUjian
	fmt.Println(lulus)
}