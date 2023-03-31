package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("In main func")
	go startFunc()
	fmt.Println("Exiting main func")
	time.Sleep(1 * time.Second)

}

func startFunc() {
	fmt.Println("In the Goroutine")
}
