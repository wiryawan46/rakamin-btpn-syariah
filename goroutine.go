package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	const timeoutDuration = 10 * time.Second

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	//ctx, cancel := context.WithTimeout(ctx, timeoutDuration)
	defer cancel()
	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			cancel()
			//return
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
				break
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}
