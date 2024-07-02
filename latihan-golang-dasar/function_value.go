package main

import "fmt"

func getGoodBye(name string) string {
	return "Hello " + name
}

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("aan"))
}
