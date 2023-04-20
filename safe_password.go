package main

import (
	"fmt"
)

func isSecurePassword(n int) bool {
	digits := make(map[int]bool)
	for n > 0 {
		digit := n % 10
		if digit < 0 || digit > 9 || digits[digit] {
			return false
		}
		digits[digit] = true
		n /= 10
	}
	return true
}

func main() {
	var start, end int
	fmt.Print("Enter start of range: ")
	fmt.Scan(&start)
	fmt.Print("Enter end of range: ")
	fmt.Scan(&end)
	fmt.Println("Secure passwords between", start, "and", end, "are:")
	for i := start; i <= end; i++ {
		if i > 9 && isSecurePassword(i) {
			fmt.Println(i)
		}
	}
}
