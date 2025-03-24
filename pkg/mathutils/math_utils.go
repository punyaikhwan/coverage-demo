package mathutils

import "fmt"

// Add adds two integers and returns the result
func Add(a, b int) int {
	return a + b
}

// Multiply multiplies two integers and returns the result
func Multiply(a, b int) int {
	return a * b
}

func CheckIsAdditionTrue(a, b int, sum int) bool {
	return Add(a, b) == sum
}

// Divide divides two integers and returns the result
// Returns an error if the divisor is zero
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
