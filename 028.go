package main

func main() {
	sum, number, skip := 1, 1, 0
	for i := 0; i < 2000; i++ {
		if i%4 == 0 {
			skip += 2
		}
		for j := 0; j < skip; j++ {
			number++
		}
		sum += number
	}
	println(sum)
}

