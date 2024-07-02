package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sumAll(1, 2, 3)
	fmt.Println("result = ", result)

	numbers := []int{10, 10, 10}
	resultSlice := sumAll(numbers...)
	fmt.Println("resultSlice = ", resultSlice)
}
