package main

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	result := 1
	var divisor int
	for i := range primes {
		divisor = primes[i]
		for divisor <= 20 {
			divisor *= primes[i]
		}
		divisor /= primes[i]
		result *= divisor
	}
	println(result)
}

