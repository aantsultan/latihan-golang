package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result any = random()

	switch value := result.(type) {
	case string:
		fmt.Println("Tipe data String", value)
	case int:
		fmt.Println("Tipe data Int", value)
	default:
		fmt.Println("Default", value)
	}
}
