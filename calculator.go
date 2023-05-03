// Package calculator does simple calculations
package calculator

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
	return a / b, nil
}
