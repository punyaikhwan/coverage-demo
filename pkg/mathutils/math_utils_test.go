package mathutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// TestMultiply tests the Multiply function
func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	expected := 6
	if result != expected {
		t.Errorf("Multiply(2, 3) = %d; want %d", result, expected)
	}
}

// TestDivide tests the Divide function
func TestDivide(t *testing.T) {
	// Test successful division
	result, err := Divide(6, 3)
	expected := 2
	if err != nil {
		t.Errorf("Divide(6, 3) returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Divide(6, 3) = %d; want %d", result, expected)
	}

	// Test division by zero
	_, err = Divide(6, 0)
	if err == nil {
		t.Errorf("Divide(6, 0) expected an error, but got nil")
	}
}

func TestCheckIsAdditionTrueTrueInput(t *testing.T) {
	result := CheckIsAdditionTrue(2, 3, 5)
	assert.True(t, result)
}

func TestCheckIsAdditionTrueFalseInput(t *testing.T) {
	result := CheckIsAdditionTrue(2, 3, 6)
	assert.False(t, result)
}
