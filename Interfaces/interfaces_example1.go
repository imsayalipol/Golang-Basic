package main
import "fmt"

type Shape interface{
	Area() int
	// Perimeter() int
}

type Rectangle struct{
	width int
	height int
}

func (r Rectangle) Area() int {
	return 2*r.width*r.height
}

func geometry(s Shape) {
	fmt.Println("s : ",s)
	fmt.Println(s.Area())
	// fmt.Println(s.Perimeter())
}

func main(){
	r := Rectangle{width:10, height:5}
	geometry(r)

}