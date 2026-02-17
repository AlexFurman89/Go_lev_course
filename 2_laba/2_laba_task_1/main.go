//## Масиви та слайси
//Створити масив `a` та слайс `b` для з десяти елементів.
//Значення елементів масиву від 1 до 10, значення елементів слайсу -  довільні.
//За допомогою циклу for знайти  суму кожного i-го елементу масиву та слайсу та
//зберегти цю суму у i-ий елемент слайсу `result`.
//За допомогою методу `fmt.Println` (або іншого довільного методу) вивести усі значення слайсу `result` на екран.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a [10]int
	b := make([]int, 10)
	result := make([]int, 10)

	for i := 0; i < 10; i++ {
		a[i] = rand.Intn(10)
		b[i] = rand.Intn(10)
	}

	for i := 0; i < 10; i++ {
		result[i] = a[i] + b[i]
	}

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("result: ", result)

}
