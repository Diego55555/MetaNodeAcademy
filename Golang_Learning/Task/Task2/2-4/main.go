package main

import "fmt"

func main() {
	rectangle := Rectangle{width: 4, height: 2}
	circle := Circle{radius: 3}

	fmt.Println("Rectangle Area", rectangle.Area())
	fmt.Println("Rectangle Perimeter", rectangle.Perimeter())
	fmt.Println("Circle Area", circle.Area())
	fmt.Println("Circle Perimeter", circle.Perimeter())
}

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	width  float32
	height float32
}

func (rectangle Rectangle) Area() float32 {
	return rectangle.width * rectangle.height
}

func (rectangle Rectangle) Perimeter() float32 {
	return (rectangle.width + rectangle.height) * 2
}

type Circle struct {
	radius float32
}

func (circle Circle) Area() float32 {
	return float32(3.14) * circle.radius * circle.radius
}

func (circle Circle) Perimeter() float32 {
	return 2 * 3.14 * circle.radius
}
