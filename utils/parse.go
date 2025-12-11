package utils

import (
	"strconv"
	"strings"
)

func ParseStringToIntegers(s string) []int {
	result := make([]int, 0)

	for _, digit := range strings.Split(s, ",") {
		number, err := strconv.Atoi(string(digit))
		if err == nil {
			result = append(result, number)
		}
	}

	return result
}
