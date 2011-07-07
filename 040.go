package main

import (
	"strconv"
)

func main() {
	i, count, limit, mul := 1, 0, 10, 1
	var tmp int
	for limit < 1000001 {
		str := strconv.Itoa(i)
		count += len(str)
		if count >= limit {
			tmp, _ = strconv.Atoi(str[len(str)-(count-limit)-1 : len(str)-(count-limit)])
			mul *= tmp
			limit *= 10
		}
		i++
	}
	println(mul)
}

