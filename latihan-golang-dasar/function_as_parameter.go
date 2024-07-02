package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

type Filter func(string) string

func sayHelloWithFilterType(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello type", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("aant", spamFilter)

	filter := spamFilter

	sayHelloWithFilter("Anjing", filter)

	sayHelloWithFilterType("aant", spamFilter)

	sayHelloWithFilterType("Anjing", filter)

}
