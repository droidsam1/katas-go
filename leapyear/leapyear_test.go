package leapyear

import "testing"

func TestLeapYearShouldBeAFunction(t *testing.T) {
	got := LeapYear(1995)

	if got == false {
		t.Fatalf("Expected LeapYear to return something")
	}

}

func LeapYear(i int) bool {
	panic("unimplemented")
}
