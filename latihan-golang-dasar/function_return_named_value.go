package main

import "fmt"

func getComplateName() (firstName string, middleName string, lastName string) {
	firstName = "aant"
	middleName = "sultan"
	lastName = "rahmanya"
	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getComplateName()
	fmt.Println(firstName, middleName, lastName)
}
