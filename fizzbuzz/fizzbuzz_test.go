package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	t.Run("should return a number when for default", func(t *testing.T) {

		input := 1

		result := FizzBuzz(input)

		expected := "1"

		if result != expected {
			t.Fatalf("expected to return the word %v", expected)
		}
	})

	t.Run("should return Fizz for multiples of three", func(t *testing.T) {

		input := 3

		result := FizzBuzz(input)

		expected := "Fizz"

		if result != expected {
			t.Fatalf("expected to return the word %v", expected)
		}
	})

	t.Run("should return Buzz for multiples of five", func(t *testing.T) {
		input := 5

		result := FizzBuzz(input)

		expected := "Buzz"

		if result != expected {
			t.Fatalf("expected to return the word %v", expected)
		}

	})
}
func assertEquals(expected, result string, t *testing.T) {
	if expected != result {
		t.Fatalf("expected to return %v instead of %v", expected, result)
	}
}
