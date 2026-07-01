package main

import (
    "testing"
)

// Test case 1: Basic addition
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {2, 3, 5},
        {-1, 1, 0},
        {0, 0, 0},
        {100, 200, 300},
    }
    for _, tt := range tests {
        result := Add(tt.a, tt.b)
        if result != tt.expected {
            t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
        }
    }
}

// Test case 2: Subtraction
func TestSubtract(t *testing.T) {
    if got := Subtract(10, 3); got != 7 {
        t.Errorf("Subtract(10, 3) = %d; want 7", got)
    }
}

// Test case 3: Multiply
func TestMultiply(t *testing.T) {
    if got := Multiply(4, 5); got != 20 {
        t.Errorf("Multiply(4, 5) = %d; want  20", got)
    }
}

// Test case 4: Division and division by zero
func TestDivide(t *testing.T) {
    result, err := Divide(10, 2)
    if err != nil || result != 5.0 {
        t.Errorf("Divide(10, 2) = %f, %v; want 5.0, nil", result, err)
    }

    _, err = Divide(10, 0)
    if err == nil {
        t.Error("Divide(10, 0) expected error, got nil")
    }
}

// Test case 5: Prime number check
func TestIsPrime(t *testing.T) {
    cases := map[int]bool{
        1:  false,
        2:  true,
        7:  true,
        9:  false,
        13: true,
    }
    for n, expected := range cases {
        if got := IsPrime(n); got != expected {
            t.Errorf("IsPrime(%d) = %v; want %v", n, got, expected)
        }
    }
}