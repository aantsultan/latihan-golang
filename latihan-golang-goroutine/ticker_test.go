package latihan_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time2 := range channel {
		fmt.Println(time2)
	}
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time2 := range ticker.C {
		fmt.Println(time2)
	}
}
