package main
import (
	"fmt"
)

func main(){
// ways of declaring slice:

var firstSlice []int
fmt.Println("firstSlice==", firstSlice)

var secondSlice = []float32 {10.11, 20.22, 30.33}
fmt.Println("secondSlice ==", secondSlice)

var thirdSlice = make([]int, 5)
fmt.Println("thirdSlice ==", thirdSlice)

var copySlice = thirdSlice
fmt.Println("Copy slice : === :", copySlice)

thirdSlice[2] = 5
fmt.Println("thirdSlice new value ===>", thirdSlice)
}