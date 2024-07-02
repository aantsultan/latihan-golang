package latihan_golang_hackerrank

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	test := "test"
	test = strings.Replace(test, "t", "x", 1)
	fmt.Println(test)
}

func TestSlice(t *testing.T) {
	test := []string{"111", "222", "333"}
	testDst := make([]string, len(test))
	copy(testDst, test)
	fmt.Println(testDst)
}

func TestStringSelect(t *testing.T) {
	test := "32111"
	fmt.Println(test[1:2])
}
