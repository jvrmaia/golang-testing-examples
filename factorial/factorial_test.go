package factorial

import (
	"fmt"
	"testing"
)

func TestFactorialRecursiveAndInteractiveMustBeEqual(t *testing.T) {
	for i := 0; i < 10; i++ {
		rec, _ := FactorialRecursive(i)
		itr, _ := FactorialInteractive(i)
		if rec != itr {
			t.Error("FactorialRecursive should be equal then FactorialInteractive")
		}
	}
}

func TestFactorialOfNegativeNumber(t *testing.T) {
	t.Run("Recursive", func(t *testing.T) {
		ans, err := FactorialRecursive(-10)
		if err == nil {
			t.Error("Factorial of negative number must be an error")
		}
		if invalidInputError.Error() != err.Error() {
			t.Error(fmt.Sprintf("\nWrong error message:\n\tExpected: %s\n\tActual: %s", invalidInputError.Error(), err.Error()))
		}
		if ans != 0 {
			t.Error("Error value must be zero and got ", ans)
		}
	})
	t.Run("Interactive", func(t *testing.T) {
		ans, err := FactorialInteractive(-1)
		if err == nil {
			t.Error("Factorial of negative number must be an error")
		}
		if invalidInputError.Error() != err.Error() {
			t.Error(fmt.Sprintf("\nWrong error message:\n\tExpected: %s\n\tActual: %s", invalidInputError.Error(), err.Error()))
		}
		if ans != 0 {
			t.Error("Error value must be zero and got ", ans)
		}
	})
}

func TestFactorialOfZero(t *testing.T) {
	t.Run("Recursive", func(t *testing.T) {
		ans, err := FactorialRecursive(0)
		if err != nil {
			t.Error(err.Error())
		}
		if ans != 1 {
			t.Error("factorial(0) should be 1")
		}
	})
	t.Run("Interactive", func(t *testing.T) {
		ans, err := FactorialInteractive(0)
		if err != nil {
			t.Error(err.Error())
		}
		if ans != 1 {
			t.Error("factorial(0) should be 1")
		}
	})
}

func TestFactorialOfIntegerGreaterThanZero(t *testing.T) {
	t.Run("Recursive", func(t *testing.T) {
		ans, err := FactorialRecursive(5)
		if err != nil {
			t.Error(err.Error())
		}
		if ans != 120 {
			t.Error("factorial(5) should be 120 and got", ans)
		}
	})
	t.Run("Interactive", func(t *testing.T) {
		ans, err := FactorialInteractive(3)
		if err != nil {
			t.Error(err.Error())
		}
		if ans != 6 {
			t.Error("factorial(3) should be 6 and got", ans)
		}
	})
}

func BenchmarkFactorial(b *testing.B) {
	b.Run("Recursive", func(b *testing.B) {
		b.ResetTimer()
		_, err := FactorialRecursive(100)
		if err != nil {
			b.Error(err.Error())
		}
	})
	b.Run("Interactive", func(b *testing.B) {
		b.ResetTimer()
		_, err := FactorialInteractive(100)
		if err != nil {
			b.Error(err.Error())
		}
	})
}
