package leapyear

func LeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0
}
