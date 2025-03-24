package mathutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	n := TwoNumber{FirstNumber: 2, SecondNumber: 3}
	result := n.Add()
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// TestMultiply tests the Multiply function
func TestMultiply(t *testing.T) {
	n := TwoNumber{FirstNumber: 2, SecondNumber: 3}
	result := n.Multiply()
	expected := 6
	if result != expected {
		t.Errorf("Multiply(2, 3) = %d; want %d", result, expected)
	}
}

// TestDivide tests the Divide function
func TestDivide(t *testing.T) {
	// Test successful division
	n := TwoNumber{FirstNumber: 6, SecondNumber: 3}
	result, err := n.Divide()
	expected := 2
	if err != nil {
		t.Errorf("Divide(6, 3) returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Divide(6, 3) = %d; want %d", result, expected)
	}

	// Test division by zero
	n = TwoNumber{FirstNumber: 6, SecondNumber: 0}
	_, err = n.Divide()
	if err == nil {
		t.Errorf("Divide(6, 0) expected an error, but got nil")
	}
}

func TestCheckIsAdditionTrueWithTrueInput(t *testing.T) {
	n := TwoNumber{FirstNumber: 2, SecondNumber: 3}
	result := n.CheckIsAdditionTrue(5)
	assert.True(t, result)
}

func TestCheckIsAdditionTrueWithFalseInput(t *testing.T) {
	n := TwoNumber{FirstNumber: 2, SecondNumber: 3}
	result := n.CheckIsAdditionTrue(6)
	assert.False(t, result)
}

// TestIsNumberPalindrome tests the IsNumberPalindrome function
func TestIsNumberPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Palindrome with single digit", 5, true},
		{"Palindrome with multiple digits", 121, true},
		{"Palindrome with even number of digits", 1221, true},
		{"Non-palindrome", 123, false},
		{"Negative number", -121, false},
		{"Zero", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsNumberPalindrome(tc.input)
			assert.Equal(t, tc.expected, result, "Input: %d", tc.input)
		})
	}
}
