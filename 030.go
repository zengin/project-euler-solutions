package main


func main() {
	limit := 9 * 9 * 9 * 9 * 9 * (5 - 1)
	sum := 0
	for i := 22; i < limit; i++ {
		if i == sumOfFifthPowers(i) {
			sum += i
		}
	}
	println(sum)
}

func sumOfFifthPowers(in int) int {
	sum := 0
	for in > 0 {
		sum += fifthPower(in % 10)
		in /= 10
	}
	return sum
}

func fifthPower(in int) int {
	return in * in * in * in * in
}

