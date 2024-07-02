package latihan_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoRoutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
func HelloWorld() {
	fmt.Println("Hai aant")
}

func TestCreateGoRoutine(t *testing.T) {
	go HelloWorld()
	fmt.Println("Ups...")

	time.Sleep(1 * time.Second)
}
