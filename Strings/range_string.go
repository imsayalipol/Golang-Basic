package main

import "fmt"

func main() {

	var name = "Taylor Swift"

	for i, v := range name {
		fmt.Printf("Index :%d Value :%s\n", i, string(v))
	}
}
