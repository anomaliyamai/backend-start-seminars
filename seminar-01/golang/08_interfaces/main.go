package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func printShapeInfo(s Shape) {
	fmt.Printf("Площадь: %.2f\n", s.Area())
	fmt.Printf("Периметр: %.2f\n", s.Perimeter())
}

func main() {
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	fmt.Println("Прямоугольник:")
	printShapeInfo(rectangle)

	fmt.Println("\nКруг:")
	printShapeInfo(circle)

	var shapes []Shape
	shapes = append(shapes, rectangle)
	shapes = append(shapes, circle)

	fmt.Println("\nВсе фигуры:")
	for i, shape := range shapes {
		fmt.Printf("Фигура %d: площадь = %.2f\n", i+1, shape.Area())
	}
}
