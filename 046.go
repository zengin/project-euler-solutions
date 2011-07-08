package main

import (
	"math"
)

func main() {
	arr := make([]bool, 1000000)
	arr[0], arr[1], arr[2] = false, false, true
	i := 3
	var composite, found bool
	var tmp, tmp2 int
loop:
	for {
		composite = false
	inner:
		for j := 2; j < i; j++ {
			if arr[j] && i%j == 0 {
				composite = true
				break inner
			}
		}
		if composite {
			found = false
		inner2:
			for j := 2; j < i; j++ {
				if arr[j] {
					tmp = (i - j)
					if tmp2 = int(math.Sqrt(float64(tmp / 2))); tmp%2 == 0 && tmp2*tmp2 == tmp/2 {
						found = true
						break inner2
					}
				}
			}
			if !found {
				println(i)
				break loop
			}
		} else { //prime
			arr[i] = true
		}
		i += 2
	}
}

