package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("tidak bisa dibagi 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println("result=", result)
	} else {
		fmt.Println("error", err.Error())
	}
}
