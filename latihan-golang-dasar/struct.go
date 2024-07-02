package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	var aant Customer
	aant.Name = "Aant Sultan"
	aant.Address = "Indonesia"
	aant.Age = 19

	fmt.Println(aant)
	fmt.Println(aant.Name)
	fmt.Println(aant.Address)
	fmt.Println(aant.Age)

	elaina := Customer{
		Name:    "elaina witch",
		Address: "japan",
		Age:     20,
	}

	fmt.Println(elaina)
	fmt.Println(elaina.Name)
	fmt.Println(elaina.Address)
	fmt.Println(elaina.Age)

	kiana := Customer{"kiana", "bulan", 20}
	fmt.Println(kiana)
	fmt.Println(kiana.Name)
	fmt.Println(kiana.Address)
	fmt.Println(kiana.Age)
}
