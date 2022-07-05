package leapyear

type Year int

func (year Year) IsLeap() bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
