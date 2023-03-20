package main
import "fmt"

type base struct{
	num int
}

type container struct{
	base
	name string
}

func (b base)describe()string{
	return fmt.Sprintf("The number is %v",b.num)
}

func main(){	

co := container{base:base{num:33}, name:"Taylor Swift"}
po := container{base{44}, "Akon"}

fmt.Println("co content is :", co.num, co.name) // co can directly access num as base in embedded in container
fmt.Println("po content is:", po.base.num, po.name)// or we can access num using base

fmt.Println("Method describe():",co.describe())    

}