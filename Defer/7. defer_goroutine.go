package main

import (
	"fmt"
	"time"
)

func main() {
	call()
	time.Sleep(2 * time.Second)
}

func call() {
	defer func() {
		go test()
	}()

	fmt.Println("Inside call()")
}

func test() {
	fmt.Println("Inside test()")
}
