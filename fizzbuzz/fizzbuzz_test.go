package fizzbuzz_test

import (
	"strconv"
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	t.Run("should return a number when for default", func(t *testing.T) {
		input := 1

		got := FizzBuzz(input)

		if got != "1" {
			t.Fatalf("expect to return the number itself")
		}
	})

}

func FizzBuzz(input int) string {
	return strconv.Itoa(input)
}
