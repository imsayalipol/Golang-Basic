package main

import "fmt"

func main() {
	//  way 1:
	myMap := make(map[string]string)
	myMap["singer"] = "Taylor Swift"
	fmt.Println("myMap is :", myMap)
	fmt.Println("value against key singer: ", myMap["singer"])

	//  way 2:
	newMap := map[int]string{1: "Taylor Swift", 2: "Dua Lipa", 3: "Ariana Grande"}
	fmt.Println("The newMap is : ", newMap)
	fmt.Println("Value against key 2:", newMap[2])

	// deleting a key pair of a map
	delete(newMap, 3)
	fmt.Println("After deleteing : ", newMap)
	// one returns the value for that index while two returns the boolean value
	// indicating if the key is present in the map or not.
	one, two := newMap[3]
	fmt.Println("one :", one)
	fmt.Println("two :", two)

	// deleting whole map : using make() craete new map and it'll delete all key value pairs of the map
	newMap = make(map[int]string)
	fmt.Println(" Deleting whole map : ", newMap)
}
