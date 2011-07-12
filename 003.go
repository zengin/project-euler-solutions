package main

func main() {
	target := int64(600851475143)
	arr := make([]bool, 10000)
	prime := 3
	var k int
	for {
		for k = 2 * prime; k < len(arr); k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < len(arr) && arr[k]; k += 2 {
		}
		if k < len(arr) {
			prime = k
			if target%int64(k) == 0 {
				target = target / int64(k)
				if target == 1 {
					println(k)
					return
				}
			}
		} else {
			break
			//prevent infinite loop in case the answer
			//is not less than 10000
		}
	}
}

