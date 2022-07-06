package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestCalculator(t *testing.T) {

	t.Run("should be implemented by a method Add", func(t *testing.T) {
		got := Add("")

		if got != 0 {
			t.Fatalf("expected to return a integer")
		}

	})
	t.Run("should accept a string wih a single number as input parameter", func(t *testing.T) {
		got := Add("1")

		if got != 1 {
			t.Fatalf("expected to return the same number")
		}
	})

	t.Run("should return the sum when a string wih a two numbers ", func(t *testing.T) {
		got := Add("1,2")

		if got != 3 {
			t.Fatalf("expected to return 3")
		}
	})
}

func Add(input string) int {
	if input == "" {
		return 0
	}

	numbers := strings.Split(input, ",")

	result := 0
	for _, n := range numbers {
		number, _ := strconv.Atoi(n)
		result += number
	}
	return result
}
