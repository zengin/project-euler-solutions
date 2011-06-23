package main

import (
	"strconv"
)

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	max, mul := 0, 0
	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			mul = i * j
			if isPalindrome(strconv.Itoa(mul)) && mul > max {
				max = mul
			}
		}
	}
	println(max)
}

