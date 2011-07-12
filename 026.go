package main

import (
	"big"
)

//using Fermat's little theorem we need to find the largest prime less than 1000
//see details: http://en.wikipedia.org/wiki/Repeating_decimal#Fractions_with_prime_denominators

func main() {
	//find primes
	arr := make([]bool, 1000)
	arr[0], arr[1] = true, true
	prime := 3
	var k int
	for {
		for k = prime * 2; k < len(arr); k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < len(arr) && arr[k]; k += 2 {
		}
		if k < len(arr) {
			prime = k
		} else {
			break
		}
	}

	//iterate over primes in reverse direction
	b := big.NewInt(0)
	var i int
	for k = prime; k > 0; k -= 2 {
		if !arr[k] && k%2 != 0 {
		inner:
			for i = 1; i < k; i++ {
				if b.Mod(b.Sub(b.Exp(big.NewInt(10), big.NewInt(int64(i)), nil), big.NewInt(1)), big.NewInt(int64(k))).Int64() == 0 {
					break inner
				}
			}
			if k-i == 1 {
				println(k)
				return
			}
		}
	}
}

