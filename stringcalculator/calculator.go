package main

import (
	"regexp"
	"strconv"
)

func Add(input string) int {
	if input == "" {
		return 0
	}

	regularExpression, _ := regexp.Compile(`\d+`)

	numbers := regularExpression.FindAllString(input, -1)

	result := 0
	for _, n := range numbers {
		number, _ := strconv.Atoi(n)
		result += number
	}
	return result
}
