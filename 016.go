package main

import (
	"strconv"
)

func main() {
	input := "2"
	for i := 0; i < 999; i++ {
		input = multipyByTwo(input)
	}
	println(sumUpDigits(input))
}

func sumUpDigits(input string) int {
	result, digit := 0, 0
	for i := range input {
		digit, _ = strconv.Atoi(string(input[i]))
		result += digit
	}
	return result
}

func multipyByTwo(input string) string {
	result, carryout, digit := "", 0, 0
	for i := len(input) - 1; i >= 0; i-- {
		digit, _ = strconv.Atoi(string(input[i]))
		digit = digit*2 + carryout
		carryout = digit / 10
		digit = digit % 10
		result = strconv.Itoa(digit) + result
	}
	if carryout > 0 {
		result = strconv.Itoa(carryout) + result
	}
	return result
}

