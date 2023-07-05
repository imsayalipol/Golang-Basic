package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	// channels of type string, array and slice
	chString := make(chan string)
	chArray := make(chan [5]int, 7)
	chSLice := make(chan []float64)

	arr := [5]int{10, 20, 30, 40, 50}
	sli := []float64{11.22, 33.44, 55.66}

	wg.Add(3)

	go StringChan(chString, &wg)
	chString <- "Golang"

	go ArrayChan(chArray, &wg)
	chArray <- arr
	close(chArray)
	go SliceChan(chSLice, &wg)
	chSLice <- sli

	wg.Wait()

}

func StringChan(chString chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	str := <-chString
	fmt.Println("String is :", str)
}

func ArrayChan(chArray chan [5]int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("length of chArray is", len(chArray))
	for v := range chArray {
		fmt.Println(v)
	}
	fmt.Println("Now length of chArray is", len(chArray))
}

func SliceChan(chSLice chan []float64, wg *sync.WaitGroup) {
	defer wg.Done()
	sli := <-chSLice
	fmt.Println("Slice is :", sli)
}
