package main

import "fmt"

type Alamat struct {
	City     string
	Province string
	Country  string
}

func main() {
	var alamat1 *Alamat = new(Alamat)
	var alamat2 *Alamat = alamat1

	alamat2.Country = "indonesia"

	fmt.Println("alamat1=", alamat1)
	fmt.Println("alamat2=", alamat2)
}
