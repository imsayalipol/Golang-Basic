package main
import( 
	"fmt"
	"reflect")

func main(){
// ways of declaring arrays
	var firstArray [3] float32
	var secondArray =[3] string {"sayali", "taylor", "ariana"}
	var thirsdArray =[...] int {1,2,3,4,5,6}
	intaArray := [2]bool {true, false}
	

	fmt.Println("Arrays are: ")
	fmt.Println(firstArray, "and Lenghth:",len(firstArray))
	fmt.Println(secondArray, "and Lenghth:",len(secondArray))
	fmt.Println(thirsdArray, "and Lenghth:",len(thirsdArray))
	fmt.Println(intaArray, "and Lenghth:",len(intaArray))

// copying an array to another array won't change the value stored in 
// a copy even if original array value changes

copyArray := secondArray
secondArray [2] = "Dua Lipa"
fmt.Println("copyArray= ", copyArray)

var slice = make([]string, 4)
fmt.Println("Type ===", reflect.TypeOf(slice))	

}


// styles of declaring Array amd Slice

// var myArray = [3]string{"First", "Second", "Third"}
// var myArray = [...]string{"First", "Second", "Third"}
// var myArray [3]string

// var mySlice = []string{"First", "Second", "Third"}
// mySlice := []string{"First", "Second", "Third"}
// var mySlice []string
// mySlice := make([]string, 3)