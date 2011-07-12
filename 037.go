package main

import (
	"strconv"
)

func main() {
	arr := make([]bool, 1000000)
	left := make([]bool, 1000000)
	right := make([]bool, 1000000)

	//manuel setting for the values less than 10
	left[2], left[3], left[5], left[7] = true, true, true, true
	right[2], right[3], right[5], right[7] = true, true, true, true
	arr[0], arr[1], arr[6], arr[9] = true, true, true, true
	a := []int{3, 5, 7}
	for i := range a {
		for k := a[i] * 2; k < len(arr); k += a[i] {
			arr[k] = true
		}
	}

	//calculate other primes and check the condition.
	var k, tmp int
	prime, counter, sum := 11, 0, 0
	for {
		if right[prime/10] {
			right[prime] = true
		}
		tmp, _ = strconv.Atoi(strconv.Itoa(prime)[1:])
		if left[tmp] {
			left[prime] = true
			if right[prime] {
				sum += prime
				counter++
				if counter == 11 {
					println(sum)
					return
				}
			}
		}
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
}

