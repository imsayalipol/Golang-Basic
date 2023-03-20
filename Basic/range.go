package main
import "fmt"

func main(){
	// iterating map with range
	m := map[int]int{1:10, 2:20, 3:30, 4:40}
	for k,v := range m{
		fmt.Println("k :",k, "v :",v)
	}
	for key := range m{					//first value always key
		fmt.Println("key:",key)
	}

	// iterating string with range gives index and Unicode code
	str := "Taylor Swift"
	for _,s := range str{
		fmt.Println("s:",s)
	}

	for i,s := range "Taylor Swift"{
		fmt.Println(i, s)
	}
}