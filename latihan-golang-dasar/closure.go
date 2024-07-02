package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		// kemampuan func() diluar block function itu sendiri
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	fmt.Println("counter = ", counter)
}
