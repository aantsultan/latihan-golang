package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	var d = c / 2
	var e = c - 2
	var f = c % 2

	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
	fmt.Println("f=", f)

	// augmented assignment
	var i = 10
	i += 10
	fmt.Println("i=", i)
	i += 5
	fmt.Println("i=", i)

	// unary operator
	var x = 10
	x++
	fmt.Println("x=", x)
	x++
	fmt.Println("x=", x)
	x--
	fmt.Println("x=", x)
	x--
	fmt.Println("x=", x)
}
