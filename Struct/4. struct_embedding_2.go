package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	name string
}

func (b base) describe() string {
	return fmt.Sprintf("The number is %v", b.num)
}

func main() {
	co := container{base: base{num: 33}, name: "Taylor Swift"} //way 1
	po := container{base{44}, "Akon"}                          // way 2
	do := container{base{num: 66}, "Dua Lipa"}                 //way 3

	fmt.Println("co :", co.num, co.name)      // co can directly access num as base is embedded in container
	fmt.Println("po :", po.base.num, po.name) // or we can access num using base
	fmt.Println("co :", do.num)

	fmt.Println("Method describe():", co.describe())

}
