package main

import (
	"big"
)

func main() {
	a, b, c, counter := big.NewInt(int64(1)), big.NewInt(int64(2)), big.NewInt(int64(0)), 2
	for len(a.String()) < 1000 {
		c.Neg(a) //c is a temp variable
		a.Add(a, b)
		b.Neg(c)
		counter++
	}
	println(counter)
}

