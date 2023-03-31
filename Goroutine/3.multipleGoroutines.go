package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("In main")
	fmt.Println(runtime.NumCPU())

	for i := 0; i < 10; i++ {
		go funcMethod(i)
	}
	fmt.Println("Exiting main")
	time.Sleep(1 * time.Second)
}

func funcMethod(a int) {
	fmt.Println(a)
}
