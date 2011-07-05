package main

import (
	"math"
)

func main() {
	var maxCounter, aInMax, bInMax int
	maxCounter = 0
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
		inner:
			for n := 0; ; n++ {
				if !isPrime(n*n + a*n + b) {
					if n > maxCounter {
						maxCounter, aInMax, bInMax = n, a, b
					}
					break inner
				}
			}
		}
	}
	println(aInMax * bInMax)
}

func isPrime(in int) bool {
	limit := int(math.Sqrt(float64(in)))
	if in < 2 || in%2 == 0 {
		return false
	} else {
		for i := 2; i < limit; i++ {
			if in%i == 0 {
				return false
			}
		}
	}
	return true
}

