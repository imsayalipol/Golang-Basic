package main

import (
	"fmt"
	// "time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func1(ch1)
	go func2(ch2)

	select {
	case msg := <-ch1:
		fmt.Println(msg)

	case ch2 <- "hey, this is for func 2":
	}

}

func func1(ch chan string) {
	ch <- "Inside func1()"
}

func func2(ch chan string) {
	// time.Sleep(1 * time.Second)
	m := <-ch
	fmt.Println(m)
}
