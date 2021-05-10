package calculator

import (
	"errors"
	"math"
)

func Add(num1, num2 float64) (float64, error) {
	return num1 + num2, nil
}

func Subtract(num1, num2 float64) (float64, error) {
	return num1 - num2, nil
}

func Multiply(num1, num2 float64) (float64, error) {
	return num1 * num2, nil
}

func Divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return math.NaN(), errors.New("Divide by Zero not defined")
	}

	return num1 / num2, nil
}
