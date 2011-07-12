package main

func main() {
	var i, j, k, ki, ij int
	result := 1.0
	for i = 1; i < 10; i++ {
		for j = 1; j < i; j++ {
			for k = 1; k < j; k++ {
				ki = 10*k + i
				ij = 10*i + j
				if ki*j == ij*k {
					result *= float64(ij) / float64(ki)
				}
			}
		}
	}
	println(int(result))
}

