package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	t.Run("should return a number when for default", func(t *testing.T) {
		input := 1

		got := FizzBuzz(input)

		if got != "1" {
			t.Fatalf("expected to return the number itself")
		}
	})

	t.Run("should return Fizz for multiples of three", func(t *testing.T) {

		input := 3

		got := FizzBuzz(input)

		if got != "Fizz" {
			t.Fatalf("expected to return the word Fizz")
		}
	})

	t.Run("should return Buzz for multiples of five", func(t *testing.T) {
		input := 5

		got := FizzBuzz(input)

		if got != "Buzz" {
			t.Fatalf("expected to resturn the word Buzz")
		}

	})

}
