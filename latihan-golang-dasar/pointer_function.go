package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "indonesia"
}

func main() {
	address := Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
