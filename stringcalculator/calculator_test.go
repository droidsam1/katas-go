package main

import "testing"

func TestCalculator(t *testing.T) {

	t.Run("should be implemented by a method Add", func(t *testing.T) {
		got := Add("")

		if got != 0 {
			t.Fatalf("expected to return a integer")
		}

	})
}

func Add(numbers string) int {
	return 0
}
