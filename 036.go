package main

import (
	"strconv"
)

func main() {
	sum := 0
	for i := 0; i < 1000000; i++ {
		if isPalindrome(strconv.Itoa(i)) {
			if isPalindrome(strconv.Itob(i, 2)) {
				sum += i
			}
		}
	}
	println(sum)
}

func isPalindrome(in string) bool {
	for i := 0; i < len(in); i++ {
		if in[i] != in[len(in)-i-1] {
			return false
		}
	}
	return true
}

