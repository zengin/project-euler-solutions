package main

import (
	"strconv"
)

var arr []int

func main() {
	arr = make([]int, 1000000)
	arr[0] = 2
	primes := 1

	var prime bool
	for i := 3; i < 1000000; i++ {
		prime = true
	inner:
		for j := 0; j < primes; j++ {
			if i%arr[j] == 0 {
				prime = false
				break inner
			}
		}
		if prime {
			arr[primes] = i
			primes++
		}
	}
	arr = arr[:primes]
	println(circularPrimes())
}

func circularPrimes() int {
	var str string
	var local int
	var found bool
	count := 0
	for i := 0; i < len(arr); i++ {
		found = true
		str = strconv.Itoa(arr[i])
	inner:
		for j := 0; j < len(str)-1; j++ {
			str = str[1:] + str[0:1]
			if local, _ = strconv.Atoi(str); !search(local, 0, len(arr)-1) {
				found = false
				break inner
			}
		}
		if found {
			count++
		}
	}
	return count
}

func search(value, low, high int) bool {
	if high < low {
		return false
	}
	mid := low + (high-low)/2
	if arr[mid] > value {
		return search(value, low, mid-1)
	} else if arr[mid] < value {
		return search(value, mid+1, high)
	}
	return true
}

