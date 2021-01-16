// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return b - a
}

// Multiply takes two numbers and returns the result of mulitipling the second
// from the first.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the second
// from the first.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad inputs %f, %f:  division by zero", a, b)
	}
	return a / b, nil
}
