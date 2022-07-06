package main

import (
	"strings"
	"testing"
)

func TestCalculator(t *testing.T) {

	t.Run("should be implemented by a method Add", func(t *testing.T) {

		got, _ := Add("")

		expected := 0

		assertEquals(expected, got, t)

	})
	t.Run("should accept a string wih a single number as input parameter", func(t *testing.T) {

		got, _ := Add("1")

		expected := 1

		assertEquals(expected, got, t)
	})

	t.Run("should return the sum when a string wih a two numbers ", func(t *testing.T) {

		got, _ := Add("1,2")

		expected := 3

		assertEquals(expected, got, t)
	})

	t.Run("should return the sum when a string wih a two numbers bigger than two digis", func(t *testing.T) {

		got, _ := Add("101,2010")

		expected := 2111

		assertEquals(expected, got, t)
	})

	t.Run("should return the sum when a string wih a three numbers ", func(t *testing.T) {

		got, _ := Add("1,2,3")

		expected := 6

		assertEquals(expected, got, t)
	})

	t.Run("should return the sum when a string wih a four numbers ", func(t *testing.T) {

		got, _ := Add("1,2,3,4")

		expected := 10

		assertEquals(expected, got, t)
	})

	t.Run("should return the sum when a string wih an unkown amount of numbers ", func(t *testing.T) {

		got, _ := Add("0,1,1,2,3,5,8,13,21,34,55,89,144")

		expected := 376

		assertEquals(expected, got, t)
	})

	t.Run("should return the sum when a string wih a two numbers with a space between ", func(t *testing.T) {

		got, _ := Add("1, 2")

		expected := 3

		assertEquals(expected, got, t)
	})

	t.Run("should return the sum when a string wih a two numbers separated by new lines ", func(t *testing.T) {

		got, _ := Add("1/n2")

		expected := 3

		assertEquals(expected, got, t)
	})

	t.Run("should return the sum when a string wih a two numbers separated by new lines and commas ", func(t *testing.T) {

		got, _ := Add("1/n2, 3")

		expected := 6

		assertEquals(expected, got, t)
	})

	t.Run("should return the sum when a string of form //[delimiter]\n[numbers…] ", func(t *testing.T) {

		got, _ := Add("//;\n1;2")

		expected := 3

		assertEquals(expected, got, t)
	})

	t.Run("should return the sum when a string of form //[delimiter]\n[numbers…] when the delimiter part is optional", func(t *testing.T) {

		got, _ := Add("\n1;2")

		expected := 3

		assertEquals(expected, got, t)
	})

	t.Run("should return an error when a negative number is present in input parameter", func(t *testing.T) {

		got, _ := Add("-1")

		expected := -1

		assertEquals(expected, got, t)
	})

	t.Run("should return an error when a negative number is present in input parameter", func(t *testing.T) {

		_, err := Add("-1")

		expected := "negatives not allowed: -1"
		assertError(expected, err, t)
	})

	t.Run("should return an error when all negative numbers present in input paramter", func(t *testing.T) {

		_, err := Add("-1,2,-3")

		expected := "negatives not allowed: -1,-3"
		assertError(expected, err, t)
	})

}

func assertError(expected string, result error, t *testing.T) {
	if !strings.Contains(expected, result.Error()) {
		t.Fatalf("expected to return %v, instead of %v", expected, result.Error())
	}
}

func assertEquals(expected, result int, t *testing.T) {

	if expected != result {
		t.Fatalf("expected to return %v, instead of %v", expected, result)
	}
}
