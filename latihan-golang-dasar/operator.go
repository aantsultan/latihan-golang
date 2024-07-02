package main

import "fmt"

type Location struct {
	Name    string
	City    string
	Country string
}

func main() {
	location1 := Location{"bekasi", "jawa barat", "indonesia"}
	location2 := &location1

	location2.City = "depok"
	fmt.Println("location1", location1)
	fmt.Println("location2", location2)

	location2 = &Location{"sungai bambu", "jakarta", "indonesia"}
	fmt.Println("location1", location1)
	fmt.Println("location2", location2)

	location3 := &location1
	*location3 = Location{"sungai bambu", "jakarta", "indonesia"}
	fmt.Println("location1", location1)
	fmt.Println("location3", location3)
}
