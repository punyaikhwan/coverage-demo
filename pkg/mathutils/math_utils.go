package mathutils

import "fmt"

type TwoNumber struct {
	FirstNumber  int
	SecondNumber int
}

// Add adds two integers and returns the result
func (a TwoNumber) Add() int {
	return a.FirstNumber + a.SecondNumber
}

// Multiply multiplies two integers and returns the result
func (a TwoNumber) Multiply() int {
	return a.FirstNumber * a.SecondNumber
}

func (a TwoNumber) CheckIsAdditionTrue(sum int) bool {
	return a.Add() == sum
}

// Divide divides two integers and returns the result
// Returns an error if the divisor is zero
func (a TwoNumber) Divide() (int, error) {
	if a.SecondNumber == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a.FirstNumber / a.SecondNumber, nil
}
