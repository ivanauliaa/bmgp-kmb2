package calculator

import "testing"

const (
	num1                   = 12
	num2                   = 4
	expectedAddition       = 16
	expectedSubstraction   = 8
	expectedMultiplication = 48
	expectedDivision       = 3
)

func TestAddition(t *testing.T) {
	if result := Addition(num1, num2); result != expectedAddition {
		t.Errorf("Expected: %d, Actual: %d", expectedAddition, result)
	}
}

func TestSubstraction(t *testing.T) {
	if result := Substraction(num1, num2); result != expectedSubstraction {
		t.Errorf("Expected: %d, Actual: %d", expectedSubstraction, result)
	}
}

func TestMultiplication(t *testing.T) {
	if result := Multiplication(num1, num2); result != expectedMultiplication {
		t.Errorf("Expected: %d, Actual: %d", expectedMultiplication, result)
	}
}

func TestDivision(t *testing.T) {
	if result := Division(num1, num2); result != expectedDivision {
		t.Errorf("Expected: %d, Actual: %d", expectedDivision, result)
	}
}
