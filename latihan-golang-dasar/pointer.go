package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	var address1 Address = Address{"bekasi", "jawa barat", "indonesia"}
	var address2 Address = address1

	address2.City = "ganti kota"

	fmt.Println("address1", address1) // tidak berubah
	fmt.Println("address2", address2) // berubah menjadi 'ganti kota'

	var address3 *Address = &address1 // & --> pointer

	address3.City = "depok"
	fmt.Println("address1", address1) // berubah menjadi depok
	fmt.Println("address3", address3) // berubah menjadi depok

}
