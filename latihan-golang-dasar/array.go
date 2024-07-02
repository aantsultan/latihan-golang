package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "aant"
	names[1] = "sultan"
	names[2] = "rahmanya"

	fmt.Println("names[0]=", names[0])
	fmt.Println("names[1]=", names[1])
	fmt.Println("names[2]=", names[2])

	var values = [3]int{
		90, 80, 95,
	}

	fmt.Println("values=", values)
	fmt.Println("values[0]=", values[0])
	fmt.Println("values[1]=", values[1])
	fmt.Println("values[2]=", values[2])

	fmt.Println("len values = ", len(values))

	values[1] = 0
	fmt.Println("values=", values)

	var values1 = [...]int{
		1, 2, 3, 5, 5, 5,
	}
	fmt.Println("values1=", values1)

}
