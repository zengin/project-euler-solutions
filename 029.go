package main

import (
	"big"
)

func main() {
	strMap := make(map[string]int)
	for a := 2; a < 101; a++ {
		for b := 2; b < 101; b++ {
			bigA := big.NewInt(int64(a))
			bigB := big.NewInt(int64(b))
			bigA.Exp(bigA, bigB, nil)
			strMap[bigA.String()] = 0
		}
	}
	println(len(strMap))
}

