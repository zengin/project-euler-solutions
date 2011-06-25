package main

import "math"

func isSquare(num int) (bool, int) {
	sqrt := int(math.Sqrt(float64(num)))
	if sqrt*sqrt == num {
		return true, sqrt
	}
	return false, -1
}

func main() {
loop:
	for a := 1; a < 500; a++ {
		for b := a; b < 500; b++ {
			result, c := isSquare(a*a + b*b)
			if result && a+b+c == 1000 {
				println(a * b * c)
				break loop
			}
		}
	}
}

