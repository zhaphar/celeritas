package celeritas

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Error("Expected 2 + 3 to equal 5")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(5, 3) != 2 {
		t.Error("Expected 5 - 3 to equal 2")
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(2, 3) != 6 {
		t.Error("Expected 2 * 3 to equal 6")
	}
}

func TestDivide(t *testing.T) {
	if Divide(6, 3) != 2 {
		t.Error("Expected 6 / 3 to equal 2")
	}
}
