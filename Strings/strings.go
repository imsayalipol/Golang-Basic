package main
import (
	"fmt" 
	"strings"
	)

func main(){
	var s = "Taylor Swift"	
	fmt.Println(s)
	fmt.Println(strings.ReplaceAll(s, "t", "O"))
	fmt.Println("Part :", s[7:12])
}