package main

import (
	"math"
)

func main() {
	sum, j := 0, 0
	for i := 1; i < 10000; i++ {
		j = sumOfProperDivisors(i)
		if i == sumOfProperDivisors(j) && i != j {
			sum += i
		}
	}
	println(sum)
}

func sumOfProperDivisors(input int) int {
	//checking until the square root is enough
	//since a divisor less than the square root corresponds
	//the other divisor greater than the square root
	limit, sum := int(math.Sqrt(float64(input))), 0
	for i := 1; i <= limit; i++ {
		if input%i == 0 {
			sum += i + (input / i)
		}
	}
	return sum - input
}

