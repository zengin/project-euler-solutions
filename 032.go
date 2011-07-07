package main

import (
	"strconv"
	"strings"
)

func main() {
	mp := map[int]int{}
	for i := 1; i < 200; i++ {
		for j := 1; j < 5000; j++ {
			if isPandigital(i, j, i*j) {
				mp[i*j] = 1
			}
		}
	}
	sum := 0
	for key, _ := range mp {
		sum += key
	}
	println(sum)
}

func isPandigital(mul1, mul2, result int) bool {
	str := strconv.Itoa(mul1) + strconv.Itoa(mul2) + strconv.Itoa(result)
	if len(str) != 9 {
		return false
	}
	for i := 1; i < 10; i++ {
		if !strings.Contains(str, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}

