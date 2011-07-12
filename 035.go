package main

import (
	"strconv"
)

func main() {
	arr := make([]bool, 1000000)
	arr[1] = true
	prime := 3
	count := 13
	var k, tmp, localCount int
	var str string
primeloop:
	for {
		for k = prime * 2; k < len(arr); k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < len(arr) && arr[k]; k += 2 {
		}
		if k < len(arr) {
			prime = k
			str = strconv.Itoa(prime)
			if prime > 100 {
				localCount = 1
				for i := 0; i < len(str)-1; i++ {
					str = str[1:] + str[0:1]
					tmp, _ = strconv.Atoi(str)
					if tmp > prime {
						continue primeloop
					} else if !arr[tmp] && tmp%2 != 0 {
						localCount++
					} else {
						continue primeloop
					}
				}
				count += localCount
			}
		} else {
			break
		}
	}
	println(count)
}

