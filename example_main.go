package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

// Thread 1
func main() {
	canal := make(chan int) // canal de comunicação entre as threads

	go publish(canal) // Thread 2

	reader(canal)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Println(x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
