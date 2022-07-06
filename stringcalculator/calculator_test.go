package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {

	t.Run("should be implemented by a method Add", func(t *testing.T) {
		got := Add("")

		assertEquals(0, got, t)

	})
	t.Run("should accept a string wih a single number as input parameter", func(t *testing.T) {
		got := Add("1")

		assertEquals(1, got, t)
	})

	t.Run("should return the sum when a string wih a two numbers ", func(t *testing.T) {
		got := Add("1,2")

		assertEquals(3, got, t)
	})
}

func assertEquals(expected, result int, t *testing.T) {

	if expected != result {
		t.Fatalf("expected to return %v, instead of %v", expected, result)
	}
}
