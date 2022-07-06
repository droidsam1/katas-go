package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	if input == "" {
		return 0, nil
	}

	negativeRegularExpression, _ := regexp.Compile(`\-d+`)
	negativeNumbers := negativeRegularExpression.FindAllString(input, -1)

	if strings.Contains(input, "-") {
		return -1, errors.New("negatives not allowed: " + strings.Join(negativeNumbers, ","))
	}

	regularExpression, _ := regexp.Compile(`\d+`)

	numbers := regularExpression.FindAllString(input, -1)

	result := 0
	for _, n := range numbers {
		number, _ := strconv.Atoi(n)
		result += number
	}
	return result, nil

}
