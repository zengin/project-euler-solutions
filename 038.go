package main

import (
	"strconv"
	"strings"
)

func main() {
	max := 0
	var tmp int
	for i := 1; i < 10000; i++ {
		str := ""
		for j := 1; j < 9 && len(str) < 9; j++ {
			str += strconv.Itoa(i * j)
		}
		if len(str) == 9 && isPandigital(str) {
			tmp, _ = strconv.Atoi(str)
			if tmp > max {
				max = tmp
			}
		}
	}
	println(max)
}

func isPandigital(str string) bool {
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

