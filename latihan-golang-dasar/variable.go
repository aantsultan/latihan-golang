package main

import "fmt"

func main() {
	// full declaration
	var name string

	name = "aant sultan"
	fmt.Println(name)

	name = "sultan r"
	fmt.Println(name)

	// no need implicit declaration
	var nameLagi = "aant singkat"
	fmt.Println(nameLagi)

	nameLagi = "sultan singkat"
	fmt.Println(nameLagi)

	// using := for init declare
	nameInit := "aant init"
	fmt.Println(nameInit)

	nameInit = "r init"
	fmt.Println(nameInit)

	// multiple var
	var (
		firstName  = "aant"
		middleName = "sultan"
		lastName   = "rahmanya"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
