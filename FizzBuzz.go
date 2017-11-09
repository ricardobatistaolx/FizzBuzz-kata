



package main

import (
	"fmt"
	"strconv"
	"errors"
)

func main() {

	for i := 1; i <= 100; i++ {
		v, _ := Say(i)
		fmt.Print(fmt.Sprintf("%s ", v))
	}
}

func Say(number int) (string, error) {
	if number < 0 {
		return "", errors.New("Less than 0")
	}

	if number % 15 == 0 {
		return "FizzBuzz", nil
	}

	if number % 3 == 0 {
		return "Fizz", nil
	}

	if number % 5 == 0 {
		return "Buzz", nil
	}

	return strconv.Itoa(number), nil
}