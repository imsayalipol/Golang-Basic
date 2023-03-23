package main

import "fmt"

func main() {

	// switch with switch statement and switch expressions
	switch ch := "a"; ch {
	case "a":
		fmt.Println("Its A")

	case "b":
		fmt.Println("Its B")

		//there can be multiple expressions in the case
	case "c", "d":
		fmt.Println("Its C or D")
	}

	//both switch statement and switch expression
	age := 45

	switch {
	case age < 24:
		fmt.Println("Teens")

	case age > 25 && age < 50:
		fmt.Println("Adults")

	case age > 50:
		fmt.Println("Old")
	}

	// Only switch statement
	switch number := 33; {
	case number < 20:
		fmt.Println("Less than 20")

	case number < 50:
		fmt.Println("Less than 50")
	}

	count := 45
	switch {
	case count < 10:
		fmt.Println("Its less than 10")
		fallthrough

	case count < 50:
		fmt.Println("Its less than 50")
		fallthrough
		// fmt.Println("Its smaller than 50")

	case count < 100:
		fmt.Println("Its less than 100")
	}
}
