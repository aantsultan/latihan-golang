package main

import "fmt"

func main() {

	name := "ts"

	if name == "aant" {
		fmt.Println("berhasil aant")
	} else if name == "sultan" {
		fmt.Println("berhasil sultan")
	} else {
		fmt.Println("gagal")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("nama panjang !")
	} else {
		fmt.Println("nama pendek")
	}

}
