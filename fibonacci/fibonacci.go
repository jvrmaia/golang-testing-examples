package fibonacci

import "fmt"

var invalidInputError = fmt.Errorf("Fibonacci input must be non-negative integer")

func FibonacciRecursive(n int) (int, error) {
	if n < 0 {
		return -1, invalidInputError
	}
	return fib(n), nil
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func FibonacciInteractive(n int) (int, error) {
	if n < 0 {
		return -1, invalidInputError
	}
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		aux := a + b
		a = b
		b = aux
	}
	return a, nil
}
