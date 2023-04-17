package main
import "fmt"

func main(){
	var a = "sayali"
	var b,c = 10,20
	var f = 10.11
	val := true
	x:=3

	fmt.Printf("b= %d, c=%d \n",b,c)
	fmt.Println("a =", a)
	fmt.Println("f=",f)
	fmt.Println("val=",val)
	fmt.Println(1<<uint(x)|1)
}

