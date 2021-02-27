package factorial

import "fmt"

var invalidInputError = fmt.Errorf("Factorial input must a non-negative integer")

func FactorialRecursive(n int) (int, error) {
	if n < 0 {
		return 0, invalidInputError
	}
	if n == 0 {
		return 1, nil
	}
	return fact(n), nil
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func FactorialInteractive(n int) (int, error) {
	if n < 0 {
		return 0, invalidInputError
	}
	acc := 1
	for i := 1; i <= n; i++ {
		acc *= i
	}
	return acc, nil
}
