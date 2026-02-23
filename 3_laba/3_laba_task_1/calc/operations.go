package calc

import "errors"

func Sum(nums ...float64) float64 {
	var sum float64 = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func Max(nums ...float64) float64 {
	var max float64 = nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
func Min(nums ...float64) float64 {
	var min float64 = nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Error: Division by zero is not possible")
	}
	return a / b, nil
}
