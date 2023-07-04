package main

import (
	"fmt"
	"sync"
)

func RunnerOne(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Name :", s)
}

func RunnerTwo(age int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Age :", age)
}

func main() {

	var wg sync.WaitGroup //or wg:=new(syn.WaitGroup)

	wg.Add(2)

	go RunnerOne("John", &wg) 
	go RunnerTwo(90, &wg)

	wg.Wait()

	fmt.Println("We exit")

	// wg.Wait()    // uncomment this, comment line 27 and run

}
