package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`aa([a-z])t`)

	fmt.Println(regex.MatchString("aant"))
	fmt.Println(regex.MatchString("aamt"))
	fmt.Println(regex.MatchString("aaNt"))

	fmt.Println(regex.FindAllString("aant aamt aart aa1t aaNt", 10))
}
