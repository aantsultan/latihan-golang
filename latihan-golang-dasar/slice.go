package main

import (
	"fmt"
)

func main() {
	// tipe slice yang penting,
	// 1. pointer (petunjuk data pertama)
	// 2. length
	// 3. capacity (kapasitas slice, dimana length tidak boleh lebih dari capacity)

	// array[low:high] -> low = 0; low < high
	// array[low:] -> low = 0; log < len

	names := [...]string{"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}

	slice1 := names[4:12]
	fmt.Println("slice1 = ", slice1)

	slice2 := names[:3]
	fmt.Println("slice2 = ", slice2)

	var slice3 []string = names[3:]
	fmt.Println("slice3 = ", slice3)

	var slice4 []string = names[:]
	fmt.Println("slice4 = ", slice4)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	fmt.Println(days)
	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)
	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "aant"
	newSlice[1] = "aant"
	// newSlice[2] = "aant" --> error, harsnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "aant")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "sultan"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	initArray := [...]int{1, 2, 3}
	initSlice := []int{1, 2, 3}

	fmt.Println(initArray)
	fmt.Println(initSlice)
}
