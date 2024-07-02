package main

import (
	"fmt"
	"latihan-golang-dasar/database"
	_ "latihan-golang-dasar/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
