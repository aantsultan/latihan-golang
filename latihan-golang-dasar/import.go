package main

import (
	"fmt"
	"latihan-golang-dasar/helper"
)

func main() {
	result := helper.SayHello("aant")
	fmt.Println("result", result)

	fmt.Println(helper.Application)
	//fmt.Println(helper.version) // tidak bisa diakses dari helper / package yang beda
	//fmt.Println(helper.sayHai) // tidak bisa diakses dari helper / package yang beda
}
