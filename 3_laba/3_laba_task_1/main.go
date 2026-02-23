package main

import (
	"3_laba_task_1/calc"
	"fmt"
)

func init() {
	fmt.Printf("Init func min: %.2f \n", calc.Min(1, 2, 3))
}

func main() {
	fmt.Println("sum result: ", calc.Sum(1, 2, 3))
	fmt.Println("max result: ", calc.Max(1, 2, 3))
	fmt.Println("min result: ", calc.Min(1, 2, 3))
	result, err := calc.Divide(10, 3)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("divide_result: %.2f", result)
	}

}
