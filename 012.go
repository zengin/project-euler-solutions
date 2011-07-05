package main

import (
	"math"
)

func main() {
	divisors := 1
	triangleN := uint64(1)
	for i := 2; divisors < 500; i++ {
		triangleN += uint64(i)
		divisors = noOfDivisors(triangleN)
	}
	println(triangleN)
}

func noOfDivisors(number uint64) int {
	counter := 0
	limit := uint64(math.Sqrt(float64(number)))
	for i := 1; uint64(i) <= limit; i++ {
		if number%uint64(i) == 0 {
			counter += 2
		}
	}
	return counter
}

