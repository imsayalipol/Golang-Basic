// package main

// import "fmt"

// func main() {
// 	x := []int{1, 2, 3, 4, 5}
// 	fmt.Println(x)

// 	fmt.Println("Deleteing elemnts 2,3 from x")
// 	x = append(x[:1], x[3:]...) // variadic func(unfurling)
// 	fmt.Println(x)

// }


package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second * 2)
}

func send(ch chan int) {
	ch <- 1
	fmt.Println("Sending value to channel complete")
}

func receive(ch chan int) {
	time.Sleep(time.Second * 1)
	fmt.Println("Timeout finished")
	//_ = <-ch
	return
}
