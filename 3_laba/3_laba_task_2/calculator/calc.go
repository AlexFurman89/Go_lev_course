package calculator

import "errors"

type Calc struct{}

func (c *Calc) Sum(nums ...float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func (c *Calc) Min(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func (c *Calc) Max(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func (c *Calc) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
