// there is no more capacity, so no new elements can be accommodated.
// So in this case under the hood an array of double the capacity will be allocated.
// The current array pointed by the  slice will be copied to that new array.
// Now the slice will starting pointing to this new array.
// Hence the capacity will be doubled and length will be increased by 1.

package main

import "fmt"

func main() {
	var mySlice = make([]int, 3, 4)
	mySlice[0] = 1
	mySlice[1] = 2
	mySlice[2] = 3

	fmt.Println("mySLice is :", mySlice)
	fmt.Println("Length :", len(mySlice))
	fmt.Println("Capacity :", cap(mySlice)) // cap is 4

	mySlice = append(mySlice, 11, 22)
	fmt.Println("After appending :", mySlice)
	fmt.Println("Length :", len(mySlice))
	fmt.Println("Capacity :", cap(mySlice)) // cap became twice the previous cap
}
