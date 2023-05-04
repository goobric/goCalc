// Package calculator does simple calculations
package calculator

import (
	"errors"
	"fmt"
	"math"
)

// func Add takes two numbers and returns the result of adding them together
func Add(a, b float64) float64 {
	return a + b
}

// func Subtract takes two numbers a & b, and returns the result of subtracting b from a
func Subtract(a, b float64) float64 {
	return a - b
}

// func Multiply takes two numbers and returns the product of the result
func Multiply(a, b float64) float64 {
	return a * b
}

// func Divide takes two numbers and returns the quotient of the result
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Sqrt func returns the square root of the input number, or an error for negative inputs.
func Sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0, fmt.Errorf("square root of negative number is not allowed: %f", input)
	}
	return math.Sqrt(input), nil
}
