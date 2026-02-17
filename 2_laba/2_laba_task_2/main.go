// - Описати інтерфейс `Shape`:
// - Area() float64
// - Perimeter() float64
// - Створити структури:
// - `Circle`
// - `Rectangle`
// - `Triangle`
// - Реалізувати обчислення площі та периметра
package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
func (t Triangle) Perimeter() float64 {
	return t.Base + t.Base + t.Height
}

func main() {

	shapeOne := Shape(Circle{Radius: 10})
	shapeTwo := Shape(Rectangle{Width: 10, Height: 20})
	shapeThree := Shape(Triangle{Base: 10, Height: 20})

	fmt.Println("ShapeOne area: ", shapeOne.Area(), "ShapeOne perimeter: ", shapeOne.Perimeter())
	fmt.Println("ShapeTwo area: ", shapeTwo.Area(), "ShapeTwo perimeter: ", shapeTwo.Perimeter())
	fmt.Println("ShapeThree area: ", shapeThree.Area(), "ShapeThree perimeter: ", shapeThree.Perimeter())
}
