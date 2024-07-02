package main

import "fmt"

func main() {
	var personbasic map[string]string = map[string]string{}
	personbasic["name"] = "aant"
	personbasic["address"] = "indo"

	person := map[string]string{
		"name":    "aant",
		"address": "indo",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "aant"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
