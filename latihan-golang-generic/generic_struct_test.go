package latihan_golang_generic

import (
	"fmt"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "aant",
		Second: "sultan",
	}

	fmt.Println(data)
}
