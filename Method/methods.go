package main
import "fmt"

type Person struct{
		Name string
		Age int
		Salary float32
	}

// accepts values i.e. copy of struct
func (p Person)name(){
	fmt.Println("The Name is", p.Name)
}

// accepts instance of struct i.e. direct access to memory address
func (p *Person)salary(){
	fmt.Println("Her salary is", p.Salary)
}

func main(){	

	singer := Person{"Taylo Swift", 33, 458125.25}
	singer.name()
	singer.salary()
	
}