package main

import (
	"3_laba_task_2/calculator"
	"fmt"
)

func main() {

	var c calculator.Calculator = &calculator.Calc{}

	fmt.Println("Sum:", c.Sum(1, 2, 3))
	fmt.Println("Max:", c.Max(1, 2, 3))
	fmt.Println("Min:", c.Min(1, 2, 3))

	result, err := c.Divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Divide: %.2f\n", result)

}
