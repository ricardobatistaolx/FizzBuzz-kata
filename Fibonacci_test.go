package main

import (
	"testing"
)

func TestShouldReturnFalseOnZero(t *testing.T) {
	if Fibonacci(0) == true {
		t.Error("0 is not a fibonacci")
	}
}

func TestShouldReturnTrueOnOne(t *testing.T) {
	if Fibonacci(1) == false {
		t.Error("1 is a fibonnacci")
	}
}

func TestShouldReturnTrueOnTwo(t *testing.T) {
	if Fibonacci(2) == false {
		t.Error("2 is a fibonnacci")
	}
}

func TestShouldReturnFalseOnFour(t *testing.T) {
	if Fibonacci(4) == true {
		t.Error("4 is not a fibonnacci")
	}
}

func TestShouldReturnTrueOnABigFibonacciNumber(t  *testing.T) {
	if Fibonacci(144) == false {
		t.Error("144 is a fibonnacci")
	}
}
