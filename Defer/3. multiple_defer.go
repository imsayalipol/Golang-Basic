package main

import "fmt"

func main() {
	fmt.Println("In main func")
	defer fmt.Println("Defer in : main()")
	f1()
	fmt.Println("Out main func")
}

func f1() {
	fmt.Println("In f1 func")
	defer fmt.Println("Defer in : f1()")
	f2()
	fmt.Println("Out f1 func")
}

func f2() {
	fmt.Println("In f2 func")
	defer fmt.Println("Defer in : f2()")
	fmt.Println("Out f2 func")
}
