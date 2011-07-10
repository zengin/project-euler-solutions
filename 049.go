package main

import (
	"strings"
	"strconv"
)

func main() {
	arr := make([]bool, 10000)
	prime := 2
	finished := false
	var i int
	for !finished {
		for i = 2 * prime; i < 10000; i += prime {
			arr[i] = true
		}
		//next prime
		for i = prime + 1; i < 10000 && arr[i]; i++ {
		}
		if i < 10000 {
			prime = i
		} else {
			finished = true
		}
	}
outer:
	for i := 1000; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if i != 1487 && !arr[i] && !arr[j] && isPerm(i, j) && 2*j-i < 10000 && !arr[2*j-i] && isPerm(i, 2*j-i) {
				print(i)
				print(j)
				print(2*j - i)
				break outer
			}
		}
	}
	println()
}

func isPerm(i1, i2 int) bool {
	s1, s2 := strconv.Itoa(i1), strconv.Itoa(i2)
	for i := 0; i < len(s1); i++ {
		//we need a cross check for repetition of numbers like in 1049 1499 comparison
		if !strings.Contains(s1, s2[i:i+1]) || !strings.Contains(s2, s1[i:i+1]) {
			return false
		}
	}
	return true
}

