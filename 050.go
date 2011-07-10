package main

var primes []int

func main() {
	arr := make([]bool, 1000000)
	arr[0], arr[1] = true, true
	count := 2
	prime := 3
	var k int
	finished := false
	for !finished {
		for k = 2 * prime; k < len(arr); k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < len(arr) && arr[k]; k += 2 {
		}
		if k < len(arr) {
			prime = k
			count++
		} else {
			finished = true
		}
	}
	primes = make([]int, count)
	index := 1
	primes[0] = 2
	for i := 0; i < len(arr); i++ {
		if !arr[i] && i%2 != 0 {
			primes[index] = i
			index++
		}
	}
	answer := 0
	maxCount := 0
	var tmp int
	for i := range primes {
		tmp = primes[i]
		count = 0
	inner2:
		for j := i + 1; j < len(primes); j++ {
			tmp += primes[j]
			count++
			if tmp > len(arr)-1 {
				break inner2
			}
			if tmp%2 != 0 && !arr[tmp] && count > maxCount { //is prime and better result
				maxCount = count
				answer = tmp
			}
		}
	}
	println(answer)
}

