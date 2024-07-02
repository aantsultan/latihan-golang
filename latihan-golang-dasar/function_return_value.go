package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	result := getHello("aant")
	fmt.Println(result)
}
