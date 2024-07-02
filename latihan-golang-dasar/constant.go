package main

import "fmt"

func main() {
	const firstName string = "aant"
	const lastName = "sultan"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		fName = "aant"
		lName = "sultan"
	)

	fmt.Println(fName)
	fmt.Println(lName)
}
