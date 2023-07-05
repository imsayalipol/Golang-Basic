package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	InterCh := make(chan interface{})

	wg.Add(1)

	go calling(&wg, InterCh)

	InterCh <- 123
	InterCh <- "sayali"
	InterCh <- 54.20
	InterCh <- true
	InterCh <- []int{10, 20, 30}

	close(InterCh)

	wg.Wait()
}

func calling(wg *sync.WaitGroup, InterCh chan interface{}) {

	defer wg.Done()

	for v := range InterCh {
		fmt.Println(v)
	}
}
