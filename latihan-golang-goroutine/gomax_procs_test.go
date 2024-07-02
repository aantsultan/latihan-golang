package latihan_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcsChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("total cpu", totalCpu)

	runtime.GOMAXPROCS(12)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine", totalGoroutine)

	group.Wait()

}

func TestGoMaxProcs(t *testing.T) {
	//group := sync.WaitGroup{}
	//group.Add(1)
	//
	//for i := 0; i < 100; i++ {
	//	go func() {
	//		time.Sleep(3 * time.Second)
	//		group.Done()
	//	}()
	//}

	totalCpu := runtime.NumCPU()
	fmt.Println("total cpu", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine", totalGoroutine)

	//group.Wait()

}
