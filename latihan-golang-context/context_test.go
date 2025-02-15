package latihan_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextWithDeadline(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounterSlow(ctx)
	fmt.Println(runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter", n)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounterSlow(ctx)
	fmt.Println(runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter", n)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}

func CreateCounterSlow(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for { // infinite looping
			select {
			case <-ctx.Done():
				return // return, akan mengagalkan goroutine
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	fmt.Println(runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("counter", n)

		if n == 10 {
			break
		}
	}
	cancel() // mengirim sinyal cancel ke context
	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for { // infinite looping
			select {
			case <-ctx.Done():
				return // return, akan mengagalkan goroutine
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b")) // nil, karena b bukan parent context F
	fmt.Println(contextA.Value("b")) // nil, tidak membaca child
}

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}
