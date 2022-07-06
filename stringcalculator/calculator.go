package main

import (
	"strconv"
	"strings"
)

func Add(input string) int {
	if input == "" {
		return 0
	}

	numbers := strings.Split(input, ",")

	result := 0
	for _, n := range numbers {
		number, _ := strconv.Atoi(strings.Trim(n, " "))
		result += number
	}
	return result
}
