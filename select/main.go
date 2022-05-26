package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genetate() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			out <- i
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d ,received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = genetate(), genetate()
	var worker = createWorker(0)
	n := 0
	var values []int
	var activeValues int
	tm := time.After(time.Second * 10)
	for {
		var activeworker chan<- int
		if len(values) > 0 {
			activeworker = worker
			activeValues = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeworker <- activeValues:
			values = values[1:]
		case <-tm:
			fmt.Println("BYE")
			return
		}
	}

}
