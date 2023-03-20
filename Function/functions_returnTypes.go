package main
import "fmt"

func calculate(a int, b int) (int, int) {
	s := a+b
	d := b-a
	return s,d	
}

func main(){

	var a,b = 10,20	
	sum, diff := calculate(a,b)
	fmt.Println("SUM =", sum, "DIFF =", diff) 
}