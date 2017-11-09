package main

func Fibonacci(number int) bool {


	if number < 1 {
		return false
	}
	if number < 3 {
		return true
	}

	n := 1
	n1 := 2

	for n + n1 < number {
		next := n + n1
		n = n1
		n1 = next
	}

	return n+n1 == number
}