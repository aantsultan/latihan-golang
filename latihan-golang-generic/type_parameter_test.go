package latihan_golang_generic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func Length[T interface{}]() {
//
//}

//func Length[T any]() {
//	// any = interface {}
//}

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

//func TestSample(t *testing.T) {
//	assert.True(t, true)
//}

func TestLength(t *testing.T) {
	var result = Length[string]("aant")
	assert.Equal(t, "aant", result)

	var resultInt = Length[int](100)
	assert.Equal(t, 100, resultInt)
}
