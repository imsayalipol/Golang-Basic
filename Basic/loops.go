package main

import "fmt"

func main() {
	fmt.Println("for loop :")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// if loop with syntax if expression; condition
	if a := 6; a > 5 {
		fmt.Println("a is greater than 5")
	}

	j := 1
	fmt.Println("for as while without condition and with break :")
	for {
		if j > 10 {
			break
		}
		fmt.Println(j)
		j++
	}

	array := [3]int{111, 222, 333}
	fmt.Println("iterarting Array :")
	for _, num := range array {
		fmt.Println(num)
	}

	slice := []float32{10.33, 20.356, 223.514}
	fmt.Println("iterating Slice :")
	for _, f := range slice {
		fmt.Println(f)
	}
	m := map[string]string{"good": "Ariana", "better": "Dua Lipa", "best": "Taylor Swift"}
	fmt.Println("Printing Map:")
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}
