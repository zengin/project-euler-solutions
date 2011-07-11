package main

func main() {
	arr := make([]bool, 105000)
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
			if count == 10001 {
				println(prime)
				break
			}
		} else {
			break
		}
	}
}

