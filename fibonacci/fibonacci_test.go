package fibonacci

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	input    int
	expected int
}

func TestNegativeInput(t *testing.T) {
	t.Run("Testing Fibonacci for negative input", func(t *testing.T) {
		t.Run("Recursive", func(t *testing.T) {
			ans, err := FibonacciRecursive(-22)
			assert.Error(t, err)
			assert.Equal(t, -1, ans)
		})
		t.Run("Interactive", func(t *testing.T) {
			ans, err := FibonacciInteractive(-1)
			assert.Error(t, err)
			assert.Equal(t, -1, ans)
		})
	})
}

func TestFibonacci(t *testing.T) {
	testData := []testData{
		testData{input: 0, expected: 0},
		testData{input: 1, expected: 1},
		testData{input: 2, expected: 1},
		testData{input: 3, expected: 2},
		testData{input: 4, expected: 3},
		testData{input: 5, expected: 5},
	}
	for _, test := range testData {
		t.Run(fmt.Sprintf("Testing Fibonacci(%d)", test.input), func(t *testing.T) {
			t.Run("Recursive", func(t *testing.T) {
				ans, err := FibonacciRecursive(test.input)
				assert.NoError(t, err)
				assert.Equal(t, test.expected, ans)
			})
			t.Run("Interactive", func(t *testing.T) {
				ans, err := FibonacciInteractive(test.input)
				assert.Nil(t, err)
				assert.Equal(t, test.expected, ans)
			})
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	b.Run("Recursive", func(b *testing.B) {
		b.ResetTimer()
		_, err := FibonacciRecursive(50)
		if err != nil {
			b.Error(err.Error())
		}
	})
	b.Run("Interactive", func(b *testing.B) {
		b.ResetTimer()
		_, err := FibonacciInteractive(50)
		if err != nil {
			b.Error(err.Error())
		}
	})
}
