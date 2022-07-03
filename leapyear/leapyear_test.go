package leapyear

import "testing"

func TestLeapYearShouldBeAFunction(t *testing.T) {
	got := LeapYear(1995)

	if got != false {
		t.Fatalf("Expected LeapYear to return a boolean")
	}
}

func TestLeapYearShouldReturnFalseForNonLeapYears(t *testing.T) {
	got := LeapYear(1995)

	if got != false {
		t.Fatalf("Expected LeapYear to return false")
	}
}

func LeapYear(i int) bool {
	return false
}
