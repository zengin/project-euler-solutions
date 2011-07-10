package main

import (
	"strconv"
	"strings"
)

func main() {
	a := make([]bool, 87654322)
	a[0], a[1] = true, true
	prime := 3
	var k int
	finished := false
	for !finished {
		for k = 2 * prime; k < len(a); k += prime {
			a[k] = true
		}
		for k = prime + 2; k < len(a) && a[k]; k += 2 {
		}
		if k < len(a) {
			prime = k
		} else {
			finished = true
		}
	}
	//a now has false values for the primes and multiples of 2
	//but we skip even numbers in our iteration.
	for i := int64(87654321); i > 0; i -= 2 {
		if !a[i] && isPandigital(i) {
			println(i)
			return
		}
	}
}

func isPandigital(in int64) bool {
	str := strconv.Itoa64(in)
	n := len(str)
	for i := 1; i <= n; i++ {
		if !strings.Contains(str, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}

