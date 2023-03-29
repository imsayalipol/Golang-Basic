package main

import "fmt"

// interface TaxCalculator
type TaxCalculator interface {
	CalculateTax() int
}

// struct India and interface method implementation
type India struct {
	CountryName   string
	TaxAmount     int
	TaxPercentage int
}

func (i India) CalculateTax() int {
	return (i.TaxAmount * i.TaxPercentage) / 100
}

// struct India and interface method implementation
type USA struct {
	CountryName   string
	TaxAmount     int
	TaxPercentage int
}

func (u USA) CalculateTax() int {
	return (u.TaxAmount * u.TaxPercentage) / 100
}

// struct India and interface method implementation
type Germany struct {
	CountryName   string
	TaxAmount     int
	TaxPercentage int
}

func (g Germany) CalculateTax() int {
	return (g.TaxAmount * g.TaxPercentage) / 100
}

func main() {

	indiaTax := India{
		CountryName:   "India",
		TaxAmount:     10000,
		TaxPercentage: 20,
	}

	usaTax := USA{
		CountryName:   "USA",
		TaxAmount:     25000,
		TaxPercentage: 15,
	}

	germanyTax := Germany{
		CountryName:   "Germany",
		TaxAmount:     30000,
		TaxPercentage: 10,
	}

	// array of type interface
	arr := []TaxCalculator{indiaTax, usaTax, germanyTax}

	// calculating tax of individual country
	for _, country := range arr {
		fmt.Println("Details :", country)
		fmt.Println("Calculated tax is :", country.CalculateTax())// runtime polymorphism

	}
}
