package main

import (
	"big"
)

func main() {
	result := big.NewInt(0)
	power := big.NewInt(0)
	var number *big.Int
	for i := 1; i < 1001; i++ {
		number = big.NewInt(int64(i))
		power.Exp(number, number, nil)
		result.Add(result, power)
	}
	str := result.String()
	println(str[len(str)-10 : len(str)])
}

