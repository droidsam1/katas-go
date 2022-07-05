package leapyear

type Year int

func LeapYear(year Year) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func (year Year) IsLeap() bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
