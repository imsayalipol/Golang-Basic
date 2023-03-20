package main

import "fmt"

type Calculate interface {
	MonthlySalary()
	AnnualSalary()
}

type Person struct {
	Name     string
	DailyPay int
}

func (p Person) MonthlySalary() {
	fmt.Println("Monthly pay :", p.DailyPay*30)

}

func (p Person) AnnualSalary() {
	fmt.Println("Yearly pay :", p.DailyPay*365)
}

func main() {

	p := Person{DailyPay: 3000}

	var c Calculate
	c = p
	fmt.Println(c)
	c.MonthlySalary()
	c.AnnualSalary()
	// amount(p)
}
