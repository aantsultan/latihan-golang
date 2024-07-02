package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAant NoKTP = "123456"
	var contoh string = "000000"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println("ktpAant = ", ktpAant)
	fmt.Println("contohKtp = ", contohKtp)
}
