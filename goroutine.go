package main

import (
	"context"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			cancel()
		}
	}

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}
