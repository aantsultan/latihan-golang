package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	// number overflow
	var nilai16 int16 = int16(nilai32)

	fmt.Println("nilai 32 = ", nilai32)
	fmt.Println("nilai 64 = ", nilai64)
	fmt.Println("nilai 16 = ", nilai16)

	var name = "Aant"
	var e = name[0]
	var eString = string(e)

	fmt.Println("name = ", name)
	fmt.Println("e = ", e)
	fmt.Println("eString = ", eString)
}
