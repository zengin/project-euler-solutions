package main

func main() {
	a, b, sum := 1, 2, 0
	for b < 4000000 {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	println(sum)
}

