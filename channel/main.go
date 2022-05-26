package main

import (
	"fmt"
	"time"
)

func CreateWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker:%d received:%c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = CreateWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 85 + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 100 + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
