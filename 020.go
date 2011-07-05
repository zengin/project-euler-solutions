package main

import (
	"big"
	"strconv"
)

func main() {
	mul := big.NewInt(1)
	for i := 2; i <= 100; i++ {
		mul.Mul(mul, big.NewInt(int64(i)))
	}
	result, digit := 0, 0
	for i := range mul.String() {
		digit, _ = strconv.Atoi(string(mul.String()[i]))
		result += digit
	}
	println(result)
}

