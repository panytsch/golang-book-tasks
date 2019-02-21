package main

import "fmt"

func main() {
	rectangle := Rectangle{2,5}
	circle := Circle{3}
	fmt.Println("Rectangle: ", rectangle, "\nCircle: ", circle)
	fmt.Println("Rectangle area: ", rectangle.area(), "\nCircle area: ", circle.area())
	fmt.Println("Rectangle perimeter: ", rectangle.perimeter(), "\nCircle perimeter: ", circle.perimeter())
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, height float64
}

func (x *Rectangle) area() float64 {
	return x.length * x.height
}

func (x *Circle) area() float64 {
	return 3.14 * x.radius * x.radius
}

func (x *Rectangle) perimeter() float64 {
	return 2 * x.length + 2 * x.height
}

func (x *Circle) perimeter() float64 {
	return 3.14 * 2 * x.radius
}