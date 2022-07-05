package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	t.Run("should return a number when for default", func(t *testing.T) {

		input := 1

		result := FizzBuzz(input)

		expected := "1"

		assertEquals(expected, result, t)
	})

	t.Run("should return Fizz for multiples of three", func(t *testing.T) {

		input := 3

		result := FizzBuzz(input)

		expected := "Fizz"

		assertEquals(expected, result, t)
	})

	t.Run("should return Buzz for multiples of five", func(t *testing.T) {
		input := 5

		result := FizzBuzz(input)

		expected := "Buzz"

		assertEquals(expected, result, t)
	})

	t.Run("should return FizzBuzz for multiples of five and three", func(t *testing.T) {
		input := 15

		result := FizzBuzz(input)

		expected := "FizzBuzz"

		assertEquals(expected, result, t)
	})
}
func assertEquals(expected, result string, t *testing.T) {
	if expected != result {
		t.Fatalf("expected to return %v instead of %v", expected, result)
	}
}
