package main

import "fmt"

func main() {

	chSend := make(chan int, 2)
	chRec := make(chan int, 2)

	sending(chSend)
	fmt.Println("chSend :", <-chSend) //note

	fmt.Println()

	chRec <- 10
	receiving(chRec)

}

func sending(chSend chan<- int) {
	chSend <- 2
}

func receiving(chRec <-chan int) {
	val := <-chRec
	fmt.Println("chRec :", val)
}
