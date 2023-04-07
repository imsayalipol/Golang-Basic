package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 3)

	go add(ch, 3)

	ch <- 2
	ch <- 1
	ch <- 2
	close(ch)

	time.Sleep(1 * time.Second)
}

func add(ch chan int, len int) {
	add := 0
	for i := 0; i < len; i++ {
		add += <-ch
	}

	// OR
	// for val := range ch {
	// 	sum += val

	fmt.Println("sum :", add)
}
