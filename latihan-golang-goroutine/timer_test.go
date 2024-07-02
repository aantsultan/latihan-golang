package latihan_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimeAfterFunction(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5*time.Second, func() { // tidak perlu membuat channel lagi
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())
	group.Wait()
}

func TestTimerAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time2 := <-channel
	fmt.Println(time2)
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time2 := <-timer.C
	fmt.Println(time2)
}
