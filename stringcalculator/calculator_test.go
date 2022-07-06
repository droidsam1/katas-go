package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {

	t.Run("should be implemented by a method Add", func(t *testing.T) {

		got := Add("")

		expected := 0

		assertEquals(expected, got, t)

	})
	t.Run("should accept a string wih a single number as input parameter", func(t *testing.T) {

		got := Add("1")

		expected := 1

		assertEquals(expected, got, t)
	})

	t.Run("should return the sum when a string wih a two numbers ", func(t *testing.T) {

		got := Add("1,2")

		expected := 3

		assertEquals(expected, got, t)
	})
}

func assertEquals(expected, result int, t *testing.T) {

	if expected != result {
		t.Fatalf("expected to return %v, instead of %v", expected, result)
	}
}
