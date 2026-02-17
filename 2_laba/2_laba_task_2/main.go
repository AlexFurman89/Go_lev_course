// ## Структури та інтерфейси
//
// ### Непарні варіанти
//
// - Описати інтерфейс `Shape`:
// - Area() float64
// - Perimeter() float64
// - Створити структури:
// - `Circle`
// - `Rectangle`
// - `Triangle`
// - Реалізувати обчислення площі та периметра
package main

func main() {

	type Shape interface {
		Area() float64
		Perimeter() float64
	}

	Circle := Shape{
		Area:=10
		Perimetr:=20
	}

}
