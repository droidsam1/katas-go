package leapyear

import "testing"

func TestLeapYear(t *testing.T) {
	t.Run("should return a boolean", func(t *testing.T) {
		got := LeapYear(1995)

		if got != false {
			t.Fatalf("Expected LeapYear to return a boolean")
		}
	})

	t.Run("should return false when parameter is a common year", func(t *testing.T) {
		got := LeapYear(1995)

		if got != false {
			t.Fatalf("Expected LeapYear to return false")
		}
	})

	t.Run("should return true for leap years", func(t *testing.T) {
		got := LeapYear(2020)

		if got != true {
			t.Errorf("Expected true, got false")
		}
	})
}
