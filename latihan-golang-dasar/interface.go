package main

import "fmt"

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

type HasName interface {
	GetName() string
}

func SayName(hashName HasName) {
	fmt.Println("hashName=", hashName.GetName())
}

func main() {
	jude := Person{"jude"}
	SayName(jude)
	kucing := Animal{"kucing"}
	SayName(kucing)
}
