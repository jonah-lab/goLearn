package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mesGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service:%s,message:%d", name, i)
			i++
		}
	}()
	return c
}
func mutiMesIn(chs ...chan string) chan string {
	c := make(chan string)
	// 如果这里运行的时候不加传参的话会触发bug,导致service1和service2无法写入数据
	for _, ch := range chs {
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}
	return c
}

func mesIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()
	go func() {
		for {
			c <- <-c2
		}
	}()
	return c
}
func funInBySelect(c1, c2 <-chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

func main() {
	m1 := mesGen("service1")
	m2 := mesGen("service2")
	m3 := mesGen("service3")
	m := mutiMesIn(m1, m2, m3)
	for {
		fmt.Println(<-m)
	}

}
