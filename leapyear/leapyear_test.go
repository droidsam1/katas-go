package leapyear

import "testing"

func TestLeapYear(t *testing.T) {
	t.Run("should return a boolean", func(t *testing.T) {
		got := Year(1995).IsLeap()

		if got != false {
			t.Fatalf("Expected isLeap() to return a boolean")
		}
	})

	t.Run("should return false when is a common year", func(t *testing.T) {
		got := Year(1995).IsLeap()

		if got != false {
			t.Fatalf("Expected isLeap() to return false")
		}
	})

	t.Run("should return true for leap years", func(t *testing.T) {
		got := Year(2020).IsLeap()

		if got != true {
			t.Errorf("Expected isLeap() to return true")
		}
	})

	t.Run("should return false for special case of non leap years", func(t *testing.T) {
		got := Year(1900).IsLeap()

		if got != false {
			t.Errorf("Expected isLeap() to return false")
		}
	})

	t.Run("should return false for special case of leap years", func(t *testing.T) {
		got := Year(1200).IsLeap()

		if got != true {
			t.Errorf("Expected isLeap() to return true")
		}
	})
}
