package main

import "fmt"

type User struct {
	Name string
}

func (user User) sayHello() {
	fmt.Println("Hello", user.Name)
}

func main() {
	jude := User{"jude"}
	jude.sayHello()
}
