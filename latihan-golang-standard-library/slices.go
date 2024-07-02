package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"aant", "sultan", "rahmanya"}
	values := []int{10,9,8}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "aant"))
	fmt.Println(slices.Index(names, "sultan"))
	fmt.Println(slices.Index(names, "rahmanya"))
}
