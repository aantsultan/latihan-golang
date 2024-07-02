package main

import "fmt"

func main() {

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("perulangan break ke ", i)
	}

	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("perulangan continue ke ", i)
	}
}
