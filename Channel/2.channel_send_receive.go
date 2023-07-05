package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	fmt.Println("Sending value to channel")
	go send(ch) //goroutine

	fmt.Println("Receiving value from channel")
	go receive(ch) //goroutine

	time.Sleep(1 * time.Second)
}

// sending value TO channel
func send(ch chan int) {
	fmt.Println("SEND")
	ch <- 1
	ch <- 10 // this won't through error
}

func receive(ch chan int) {
	fmt.Println("REC")
	val := <-ch
	fmt.Println("Received Value is :", val)
}
