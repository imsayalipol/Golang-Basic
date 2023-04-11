package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go function1(ch1)
	go function2(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)

	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}

func function1(ch chan string) {
	ch <- "Function 1 returning"

}

func function2(ch chan string) {
	ch <- "Function 1 returning"
}
