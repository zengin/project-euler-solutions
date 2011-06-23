package main

func main() {
	i, sum := 0, 0
	for i < 1000 {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
		i++
	}
	println(sum)
}

