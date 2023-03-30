package main

import (
    "fmt"
    "math"
)

// Define an interface called Shape
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Define a Circle type that implements the Shape interface
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Define a Rectangle type that implements the Shape interface
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2*r.Width + 2*r.Height
}

func main() {
    // Create a slice of Shape interfaces, and add a Circle and Rectangle to it
    shapes := []Shape{
        Circle{Radius: 5.0},
        Rectangle{Width: 10.0, Height: 5.0},
    }

    // Loop through the shapes and print out their areas and perimeters
    for _, shape := range shapes {
        fmt.Printf("Area: %f, Perimeter: %f\n", shape.Area(), shape.Perimeter())
    }
}