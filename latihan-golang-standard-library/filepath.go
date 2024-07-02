package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// package filepath ditentukan oleh sistem operasi,
	// contoh windows menggunakan backslash \
	// sedangkan non-windows seperti linux menggunakan slash /
	fmt.Println(filepath.Dir("hello\\world.go"))
	fmt.Println(filepath.Base("hello\\world.go"))
	fmt.Println(filepath.Ext("hello\\world.go"))
	fmt.Println(filepath.IsAbs("hello\\world.go")) // true jika diambil dari root folder
	fmt.Println(filepath.IsLocal("hello\\world.go"))
	fmt.Println(filepath.Join("hello","world", "main.go"))

}
