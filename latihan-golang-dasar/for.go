package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	fmt.Println("Selesai")

	for count := 1; count <= 10; count++ {
		fmt.Println("Perulangan baru ke ", count)
	}
	fmt.Println("Selesai")

	names := []string{"aant", "sultan", "rahmanya"}
	for i := 0; i < len(names); i++ {
		fmt.Println("nama : ", names[i])
	}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println("_ index ", name)
	}
}
