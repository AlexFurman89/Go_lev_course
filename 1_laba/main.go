//https://www.notion.so/1-Golang-304f115e47eb81daade2d39ccb1ee514?source=copy_link
//Завдання. Написати програму для виводу на екран результату розрахунку  функцій виду $y = f(x)$
//залежно від значення випадкового параметру $x$  та з урахуванням умов, коли має виконуватися розрахунок
//Функції 1 або Функції 2
//Необхідно реалізувати вирішення завдання із застосуванням конструкцій `if … else` та `switch … case` .

//https://www.notion.so/1-Golang-304f115e47eb81daade2d39ccb1ee514?source=copy_link
//Завдання. Написати програму для виводу на екран результату розрахунку  функцій виду $y = f(x)$
//залежно від значення випадкового параметру $x$  та з урахуванням умов, коли має виконуватися розрахунок
//Функції 1 або Функції 2
//Необхідно реалізувати вирішення завдання із застосуванням конструкцій `if … else` та `switch … case` .

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func ex1_if(x int32) float64 {
	if x > 100 {
		return float64((2*x + 3) * (4 - x))
	}
	return (float64(x) - 4) / float64(x)
}

func ex1_switch(x int32) float64 {
	switch {
	case x > 100:
		return float64((2*x + 3) * (4 - x))
	default:
		return (float64(x) - 4) / float64(x)
	}
}

func ex2_if(x int32) float64 {
	if x >= 1 && x <= 50 {
		return (float64(x) - 4) / float64(x)
	}
	return math.Pow(float64(x), 2) / (4 + math.Pow(float64(x), 2))
}

func ex2_switch(x int32) float64 {
	switch {
	case x >= 1 && x <= 50:
		return (float64(x) - 4) / float64(x)
	default:
		return math.Pow(float64(x), 2) / (4 + math.Pow(float64(x), 2))
	}
}

func ex3_if(x int32) float64 {
	if x >= 30 && x <= 60 {
		return float64(x) + 2 - math.Pow(float64(x), 3)
	}
	return math.Pow(float64(x), 2) / (4 + math.Pow(float64(x), 2))
}

func ex3_switch(x int32) float64 {
	switch {
	case x >= 30 && x <= 60:
		return float64(x) + 2 - math.Pow(float64(x), 3)
	default:
		return math.Pow(float64(x), 2) / (4 + math.Pow(float64(x), 2))
	}
}

func ex4_if(x int32) float64 {
	if x > 30 {
		return float64(x) + 2 - math.Pow(float64(x), 3)
	}
	return math.Pow(float64(x), 3) - math.Pow(float64(x), 2) - float64(x)
}

func ex4_switch(x int32) float64 {
	switch {
	case x > 30:
		return float64(x) + 2 - math.Pow(float64(x), 3)
	default:
		return math.Pow(float64(x), 3) - math.Pow(float64(x), 2) - float64(x)
	}
}

func main() {
	x := rand.Int31n(1000) // [0, 1000)

	fmt.Println("Random x =", x)

	var condition int
	var mode int

	fmt.Print("Enter condition number (1-4): ")
	fmt.Scan(&condition)

	fmt.Print("Enter mode: 1 (if/else) or 2 (switch): ")
	fmt.Scan(&mode)

	switch condition {

	case 1:
		if mode == 1 {
			fmt.Println("Result:", ex1_if(x))
		} else {
			fmt.Println("Result:", ex1_switch(x))
		}

	case 2:
		if mode == 1 {
			fmt.Println("Result:", ex2_if(x))
		} else {
			fmt.Println("Result:", ex2_switch(x))
		}

	case 3:
		if mode == 1 {
			fmt.Println("Result:", ex3_if(x))
		} else {
			fmt.Println("Result:", ex3_switch(x))
		}

	case 4:
		if mode == 1 {
			fmt.Println("Result:", ex4_if(x))
		} else {
			fmt.Println("Result:", ex4_switch(x))
		}
	default:
		fmt.Println("Invalid condition number")
	}
}
