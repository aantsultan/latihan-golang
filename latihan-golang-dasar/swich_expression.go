package main

import "fmt"

func main() {

	name := "Aant Sultan"

	switch name {
	case "Aant":
		fmt.Println("Hello aant")
	case "Sultan":
		fmt.Println("Hello Sultan")
	default:
		fmt.Println("Hello default")
	}

	switch length := len(name); length < 5 {
	case true:
		fmt.Println("BENAR !")
	case false:
		fmt.Println("SALAH !")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("> 10 !")
	case length > 5:
		fmt.Println("> 5")
	default:
		fmt.Println("default")
	}
}
