package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	aant := Man{"aant"}
	aant.Married()

	fmt.Println(aant)
}
