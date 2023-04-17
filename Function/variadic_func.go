package main

import "fmt"

func variadic(numbers ...float32) float32 {
	var sum float32
	for _, n := range numbers {
		sum += n
	}

	return sum
}
func main() {
	total := variadic(1, 2, 3, 4, 5, 6)
	fmt.Println(" Sum is : ", total)

	sum := variadic(11, 22, 333)
	fmt.Println("Another sum is :", sum)
}
