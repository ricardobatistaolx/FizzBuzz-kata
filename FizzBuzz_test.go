


package main

import (
	"testing"
)


func TestNumberIsBiggerThanZero(t *testing.T) {

	_, error := Say(-1)
	if error == nil {
		t.Errorf("should have returned error")
	}

}

func TestShouldReturnTheNumberAsString(t *testing.T) {
	value, _ := Say(2)
	if (value != "2") {
		t.Errorf("%f should have been '2'", value)
	}
}

func TestShouldReturnFizzOnDivisibleBy3(t *testing.T) {

	value, _ := Say(18)
	if (value != "Fizz") {
		t.Errorf("18 should have been 'Fizz', but was %s", value)
	}

}

func TestShouldReturnBuzzOnDivisibleBy5(t *testing.T) {

	value, _ := Say(10)
	if (value != "Buzz") {
		t.Errorf("10 should have been 'Buzz', but was %s", value)
	}
}

func TestShoulReturnFizzBuzzOnDivisibleBy3and5(t *testing.T) {

	value, _ := Say( 30)
	if (value != "FizzBuzz") {
		t.Errorf("30 should have been 'FizzBuzz', but was %s", value)
	}
}

func TestShoulReturnFizzBuzzOnDivisibleBy3and5_2(t *testing.T) {

	value, _ := Say( 15)
	if (value != "FizzBuzz") {
		t.Errorf("15 should have been 'FizzBuzz', but was %s", value)
	}
}