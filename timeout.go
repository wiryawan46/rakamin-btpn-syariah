package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	go doSomething(ch)

	select {
	case <-ctxTimeout.Done():
		fmt.Printf("Context cancelled: %v\n", ctxTimeout.Err())
	case result := <-ch:
		fmt.Printf("Received: %s\n", result)
	}
}

func doSomething(ch chan string) {
	fmt.Println("doSomething Sleeping...")
	time.Sleep(time.Second * 6)
	fmt.Println("doSomething Wake up...")
	ch <- "Did Something"
}
