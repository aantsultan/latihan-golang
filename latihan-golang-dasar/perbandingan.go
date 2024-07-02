package main

import "fmt"

func main() {
	const resultString = "result="
	var name1 = "aant"
	var name2 = "aant"

	var result bool = name1 == name2

	fmt.Println(resultString, result)

	result = name1 != name2
	fmt.Println(resultString, result)

	name1 = "a"
	name2 = "b"
	result = name1 > name2
	fmt.Println(resultString, result)

}
