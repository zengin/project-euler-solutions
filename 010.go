package main

func main() {
	arr := make([]bool, 2000000)
	arr[0], arr[1] = true, true
	sum, prime := int64(5), 3
	var k int
	for {
		for k = 2 * prime; k < len(arr); k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < len(arr) && arr[k]; k += 2 {
		}
		if k < len(arr) {
			prime = k
			sum += int64(k)
		} else {
			break
		}
	}
	println(sum)
}

