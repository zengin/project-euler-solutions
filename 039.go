package main

func main() {
	answer := 0
	max := 0
	var solutions int
	for p := 5; p <= 1000; p++ {
		solutions = 0
		for c := p; c > p/3; c-- {
			for b := 1; b < c; b++ {
				a := p - c - b
				if isPythagoras(a, b, c) {
					solutions++
				}
			}
		}
		if solutions > max {
			max, answer = solutions, p
		}
	}
	println(answer)
}

func isPythagoras(a, b, c int) bool {
	if a*a+b*b == c*c {
		return true
	}
	return false
}

