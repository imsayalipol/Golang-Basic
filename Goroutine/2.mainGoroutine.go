package main

import (
	"fmt"
	"time"
)

func main() {
	go func1()
	fmt.Println("in main")
	time.Sleep(1* time.Second)
	fmt.Println("last line in main func")
}

func func1(){
	go func2()
	fmt.Println("inside func 1")
}

func func2(){
	fmt.Println("inside func 2")
}