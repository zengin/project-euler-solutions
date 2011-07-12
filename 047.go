package main

func main() {
	arr := make([]bool, 1000000)
	arr[0], arr[1] = true, true
	count, prime := 2, 3
	var k int
	for {
		for k = 2 * prime; k < len(arr); k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < len(arr) && arr[k]; k += 2 {
		}
		if k < len(arr) {
			prime = k
			count++
		} else {
			break
		}
	}
	primes := make([]int, count)
	primes[0] = 2
	index := 1
	for i := 3; i < len(arr); i += 2 {
		if !arr[i] {
			primes[index] = i
			index++
		}
	}
	var consecutiveCount, divisorCount, tmp int
outer:
	for i := 646; i < 1000000; i++ {
		consecutiveCount = 0
	inner:
		for j := i; j < i+4; j++ {
			divisorCount = 0
			//find divisors of j
			if !arr[j] && j%2 != 0 {
				//a quick check: if j is prime
				continue outer
			} else {
				tmp = j
				for k := 0; k < len(primes); k++ {
					if (!arr[tmp] && tmp%2 != 0) || tmp%primes[k] == 0 {
						divisorCount++
						if divisorCount == 4 {
							consecutiveCount++
							if consecutiveCount == 4 {
								println(i)
								return
							}
							continue inner
						} else if !arr[tmp] && tmp%2 != 0 {
							//another quick check for prime tmp
							continue outer
						}
						for tmp%primes[k] == 0 {
							tmp /= primes[k]
						}
					}
				}
			}
		}
	}
}

